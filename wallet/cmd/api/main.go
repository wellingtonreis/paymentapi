package main

import (
	"context"
	"log"
	config "wallet/internal/configs"
	handler "wallet/internal/handlers"
	service "wallet/internal/services"

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
		if msg, err := service.Notification(ctx); err != nil {
			log.Printf("Erro no consumidor: %v", err)
		} else {
			log.Println(msg)
		}
	}()

	// time.Sleep(180 * time.Second)
	// cancel()

	app.Get("/", handler.Run)
	app.Listen(":3003")
}
