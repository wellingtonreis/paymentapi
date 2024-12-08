package services

import (
	rabbitmq "cashback/pkg/rabbitmq"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	domain "cashback/internal/domain"
	dto "cashback/internal/dto"

	viper "github.com/spf13/viper"
)

func EventTransferNotification(ctx context.Context) (string, error) {

	url := viper.GetString("AMQP.URL")
	exchange := viper.GetString("AMQP.EXCHANGE")

	queue_payment_success := viper.GetString("AMQP.QUEUE.Q1")
	queue_payment_cashback := viper.GetString("AMQP.QUEUE.Q2")

	ch, err := rabbitmq.OpenChannel(url)
	if err != nil {
		return "", fmt.Errorf("erro ao abrir o canal do RabbitMQ: %w", err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)
	go func() {
		if err := rabbitmq.Consume(
			ch,
			msgs,
			exchange,
			"topic",
			queue_payment_success,
			queue_payment_success,
		); err != nil {
			log.Printf("Erro no consumidor: %v", err)
		}
	}()

	for {
		select {
		case <-ctx.Done():

			return "Consumidor finalizado!", nil
		case msg := <-msgs:

			if msg.Body == nil {
				log.Fatalf("Mensagem vazia recebida: %v", msg)
			}

			msg.Ack(false)
			var notification dto.Notification
			if err := json.Unmarshal(msg.Body, &notification); err != nil {
				log.Fatalf("Erro ao desserializar a mensagem JSON: %v", err)
			}

			cashback, err := domain.NewCashback(notification.Amount, 0.02)
			if err != nil {
				log.Fatalf("Erro ao criar o cashback: %v", err)
			}

			cashback_balance := cashback.Calculate()
			walletsJSON, err := json.Marshal(domain.Wallet{
				UserID:   notification.UserID,
				OrderID:  notification.OrderID,
				Cashback: cashback_balance,
			})

			if err != nil {
				log.Fatalf("Erro ao serializar a batch de wallets: %v", err)
			}

			log.Printf("Mensagem enviada: %v", string(walletsJSON))

			err = rabbitmq.Publish(
				ch,
				string(walletsJSON),
				exchange,
				queue_payment_cashback,
			)
			if err != nil {
				log.Fatalf("Erro ao publicar a mensagem: %v", err)
			}

		case <-time.After(10 * time.Second):
			log.Println("Aguardando mensagens...")
		}
	}
}
