package main

import (
	config "cashback/internal/configs"
	handler "cashback/internal/handlers"
	service "cashback/internal/services"
	"context"
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func init() {
	config.Setup()
}

func main() {
	app := fiber.New()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if msg, err := service.EventTransferNotification(ctx); err != nil {
			log.Printf("Erro no consumidor: %v", err)
		} else {
			log.Println(msg)
		}
	}()

	// time.Sleep(180 * time.Second)
	// cancel()

	app.Get("/", handler.Run)
	app.Listen(":3002")
}
