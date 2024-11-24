package services

type GatewayPaymentInterface interface {
	ProcessPayment() (string, error)
}

type PaymentService struct {
	gateway GatewayPaymentInterface
}

func NewGatewayPaymentService(gateway GatewayPaymentInterface) *PaymentService {
	return &PaymentService{gateway: gateway}
}

func (p *PaymentService) Pay() (string, error) {
	return p.gateway.ProcessPayment()
}
