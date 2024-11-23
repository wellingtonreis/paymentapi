package di

import (
	config "payments/internal/configs"
	handler "payments/internal/handlers"
	repository "payments/internal/repositories"
	usecase "payments/internal/usecases"
)

type ContainerPayments struct {
	Client            *config.MongoDB
	PaymentRepository *repository.PaymentRepository
	PaymentUseCase    *usecase.PaymentUseCase
	PaymentHandler    *handler.PaymentHandler
}

func BuildContainerPayments() (*ContainerPayments, error) {
	client, _ := config.ConnectDB()
	paymentRepository := repository.NewPaymentRepository(client)
	paymentUseCase := usecase.NewPaymentUseCase(paymentRepository)
	paymentHandler := handler.NewPaymentHandler(paymentUseCase)

	return &ContainerPayments{
		Client:            client,
		PaymentRepository: paymentRepository,
		PaymentUseCase:    &paymentUseCase,
		PaymentHandler:    &paymentHandler,
	}, nil
}
