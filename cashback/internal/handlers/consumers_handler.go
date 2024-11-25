package handlers

import (
	rabbitmq "cashback/pkg/rabbitmq"
	"encoding/json"
	"fmt"
	"log"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"

	domain "cashback/internal/domain"
	dto "cashback/internal/dto"

	viper "github.com/spf13/viper"
)

func Consumer(ctx *fiber.Ctx) error {

	url := viper.GetString("AMQP.URL")
	queue := viper.GetString("AMQP.QUEUE")
	exchange := viper.GetString("AMQP.EXCHANGE")

	ch, err := rabbitmq.OpenChannel(url)
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgs, exchange, "topic", queue, "payment.success")

	for {
		select {
		case msg := <-msgs:
			fmt.Printf("Nova mensagem recebida: %s\n", msg.Body)
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
			fmt.Printf("Cashback calculado: %v\n", cashback_balance)
			// services.Save(&notification)

		case <-time.After(time.Second * 2):
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Consumidor iniciado!",
			})
		}
	}

}
