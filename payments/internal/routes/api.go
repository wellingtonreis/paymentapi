package routes

import (
	"payments/internal/di"

	fiber "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	container, err := di.BuildContainerPayments()
	if err != nil {
		panic(err)
	}

	app.Post("/make-payment", container.PaymentHandler.CreatePayment)
}
