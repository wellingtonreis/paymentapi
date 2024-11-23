package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	servicePayment := api.Group("/service")
	servicePayment.Post("/payment", func(c *fiber.Ctx) error {

		err := proxy.Do(c, "http://payments:3001/make-payment")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return nil
	})
}
