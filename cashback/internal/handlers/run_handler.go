package handlers

import (
	fiber "github.com/gofiber/fiber/v2"
)

func Run(ctx *fiber.Ctx) error {

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "*** Microservice started ***"})
}
