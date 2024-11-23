package main

import (
	config "apigateway/internal/configs"
	"apigateway/internal/routes"

	viper "github.com/spf13/viper"

	fiber "github.com/gofiber/fiber/v2"
	cors "github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	config.Setup()
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	whitelist := viper.GetString("WHITE_LIST")
	app.Use(cors.New(cors.Config{
		AllowOrigins:     whitelist,
		AllowMethods:     "POST",
		AllowHeaders:     "Accept, Authorization, Content-Type, X-CSRF-Token",
		ExposeHeaders:    "Link",
		AllowCredentials: false,
		MaxAge:           300,
	}))

	routes.SetupRoutes(app)
	app.Listen(":3000")
}
