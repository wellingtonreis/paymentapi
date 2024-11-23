package handlers

import (
	"fmt"
	dto "payments/internal/dto"
	usecases "payments/internal/usecases"

	validator "github.com/go-playground/validator/v10"
	fiber "github.com/gofiber/fiber/v2"
)

var validatePayment = validator.New()

type PaymentHandler struct {
	paymentUseCase usecases.PaymentUseCase
}

func NewPaymentHandler(uc usecases.PaymentUseCase) PaymentHandler {
	return PaymentHandler{paymentUseCase: uc}
}

func (h *PaymentHandler) CreatePayment(c *fiber.Ctx) error {

	var input dto.PaymentDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	err := validatePayment.Struct(input)
	if err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("O campo %s é inválido: %s", err.Field(), err.Tag())
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
	}

	newPayment, err := h.paymentUseCase.CreatePayment(&input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(newPayment)
}
