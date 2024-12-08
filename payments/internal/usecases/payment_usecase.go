package usecases

import (
	"context"
	"encoding/json"
	domain "payments/internal/domain"
	dto "payments/internal/dto"
	service "payments/internal/services"

	"time"
)

type PaymentRepositoryInterface interface {
	Save(ctx context.Context, payment *domain.Order) (string, error)
	Update(ctx context.Context, payment *domain.Order) error
}

type PaymentUseCase struct {
	repository PaymentRepositoryInterface
}

func NewPaymentUseCase(repository PaymentRepositoryInterface) PaymentUseCase {
	return PaymentUseCase{repository: repository}
}

func (u *PaymentUseCase) CreatePayment(input *dto.PaymentDTO) (string, error) {

	order, err := domain.NewOrder(
		input.Order.UserID,
		input.Order.OrderID,
		input.Order.CustomerName,
		domain.ShippingAddress{
			Street:     input.Order.ShippingAddress.Street,
			City:       input.Order.ShippingAddress.City,
			State:      input.Order.ShippingAddress.State,
			Country:    input.Order.ShippingAddress.Country,
			PostalCode: input.Order.ShippingAddress.PostalCode,
		},
	)
	if err != nil {
		return "", err
	}

	for _, product := range input.Product {
		order.AddProduct(domain.Product{
			ProductID:   product.ProductID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
		})
	}

	order.CalculateTotal()
	order.UpdateStatus("pending")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := u.repository.Save(ctx, order)
	if err != nil {
		return "", err
	}

	var gateway service.GatewayPaymentInterface
	gateway = &service.Wiremock{
		Amount: order.Total,
		Method: input.Payment.Method,
	}

	PaymentService := service.NewGatewayPaymentService(gateway)
	_, err = PaymentService.Pay()
	if err != nil {
		return "", err
	}

	order.UpdateStatus("shipped")
	paymentNotification := &domain.PaymentNotification{
		UserID:        order.UserID,
		OrderID:       order.OrderID,
		Amount:        order.Total,
		PaymentStatus: order.Status,
	}

	notification, err := json.Marshal(paymentNotification)
	if err != nil {
		return "", err
	}

	_, err = service.SendMessage(string(notification))
	if err != nil {
		return "", err
	}

	err = u.repository.Update(ctx, order)
	if err != nil {
		return "", err
	}

	return id, nil
}
