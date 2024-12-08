package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	rabbitmq "wallet/pkg/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"

	config "wallet/internal/configs"
	domain "wallet/internal/domain"
	dto "wallet/internal/dto"
	repositories "wallet/internal/repositories"

	viper "github.com/spf13/viper"
)

func Notification(ctx context.Context) (string, error) {

	url := viper.GetString("AMQP.URL")
	exchange := viper.GetString("AMQP.EXCHANGE")

	queue_payment_cashback := viper.GetString("AMQP.QUEUE.Q1")

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
			queue_payment_cashback,
			queue_payment_cashback,
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

			log.Printf("Mensagem recebida: %v", notification)

			client, _ := config.ConnectDB()
			repositories.NewWalletRepository(client).Update(ctx, &domain.Wallet{
				UserID:  notification.UserID,
				OrderID: notification.OrderID,
				Amount:  notification.Amount,
			})

		case <-time.After(10 * time.Second):
			log.Println("Aguardando mensagens...")
		}
	}
}
