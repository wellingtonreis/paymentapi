package dto

import (
	"time"
)

type Payment struct {
	PaymentID       string      `json:"payment_id" validate:"required"`
	Amount          float64     `json:"amount" validate:"required,gt=0"`
	Currency        string      `json:"currency" validate:"required,len=3"`
	Method          string      `json:"method" validate:"required,oneof=credit_card paypal"`
	Status          string      `json:"status" validate:"required,oneof=completed pending failed"`
	TransactionDate time.Time   `json:"transaction_date" validate:"required"`
	CardDetails     CardDetails `json:"card_details" validate:"required"`
}

type CardDetails struct {
	CardHolder     string `json:"card_holder" validate:"required"`
	CardLastDigits string `json:"card_last_digits" validate:"required,len=4,numeric"`
	ExpiryDate     string `json:"expiry_date" validate:"required"`
}

type Order struct {
	OrderID         string          `json:"order_id" validate:"required"`
	CustomerName    string          `json:"customer_name" validate:"required"`
	ShippingAddress ShippingAddress `json:"shipping_address" validate:"required"`
	OrderDate       time.Time       `json:"order_date" validate:"required"`
	Status          string          `json:"status" validate:"required,oneof=shipped pending cancelled"`
}

type ShippingAddress struct {
	Street     string `json:"street" validate:"required"`
	City       string `json:"city" validate:"required"`
	State      string `json:"state" validate:"required"`
	Country    string `json:"country" validate:"required"`
	PostalCode string `json:"postal_code" validate:"required,numeric,len=5"`
}

type Product struct {
	ProductID   string  `json:"product_id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Quantity    int     `json:"quantity" validate:"required,gt=0"`
	Category    string  `json:"category" validate:"required"`
}

type PaymentDTO struct {
	Payment Payment   `json:"payment" validate:"required"`
	Order   Order     `json:"order" validate:"required"`
	Product []Product `json:"product" validate:"required"`
}
