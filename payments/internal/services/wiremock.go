package services

import (
	"encoding/json"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
)

type Wiremock struct {
	Amount float64 `json:"amount"`
	Method string  `json:"payment_method"`
}

func (w *Wiremock) ProcessPayment() (string, error) {
	agent := fiber.Post("http://wiremock:8080/api/payments")
	agent.Set("Content-Type", "application/json")

	bytes, err := json.Marshal(w)
	if err != nil {
		return "", err
	}

	agent.Body(bytes)
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return "", errs[0]
	}

	if statusCode != fiber.StatusOK && statusCode != fiber.StatusCreated {
		return "", fmt.Errorf("unexpected status code: %d", statusCode)
	}

	jsonString := string(body)
	return jsonString, nil
}
