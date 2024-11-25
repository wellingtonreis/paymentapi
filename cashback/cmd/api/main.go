package main

import (
	config "cashback/internal/configs"
	handler "cashback/internal/handlers"

	fiber "github.com/gofiber/fiber/v2"
)

func init() {
	config.Setup()
}

func main() {
	app := fiber.New()

	app.Get("/", handler.Consumer)

	app.Listen(":3002")
}
