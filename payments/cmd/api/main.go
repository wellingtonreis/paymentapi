package main

import (
	config "payments/internal/configs"
	"payments/internal/routes"

	fiber "github.com/gofiber/fiber/v2"
)

func init() {
	config.Setup()
}

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3001")
}
