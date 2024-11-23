package domain

import (
	"errors"
	"time"
)

type ShippingAddress struct {
	Street     string `bson:"street"`
	City       string `bson:"city"`
	State      string `bson:"state"`
	Country    string `bson:"country"`
	PostalCode string `bson:"postal_code"`
}

type Order struct {
	OrderID         string          `bson:"order_id"`
	CustomerName    string          `bson:"customer_name"`
	ShippingAddress ShippingAddress `bson:"shipping_address"`
	OrderDate       time.Time       `bson:"order_date"`
	Status          string          `bson:"status"`
	Products        []Product       `bson:"products"`
	total           float64         `bson:"total"`
}

type Product struct {
	ProductID   string  `bson:"product_id"`
	Name        string  `bson:"name"`
	Description string  `bson:"description"`
	Price       float64 `bson:"price"`
	Quantity    int     `bson:"quantity"`
}

func NewOrder(orderID, customerName string, address ShippingAddress) (*Order, error) {
	if orderID == "" {
		return nil, errors.New("orderID is required")
	}
	if customerName == "" {
		return nil, errors.New("customerName is required")
	}

	return &Order{
		OrderID:         orderID,
		CustomerName:    customerName,
		ShippingAddress: address,
		OrderDate:       time.Now(),
		Status:          "pending",
		Products:        []Product{},
		total:           0.0,
	}, nil
}

func (o *Order) AddProduct(product Product) {
	o.Products = append(o.Products, product)
}

func (o *Order) CalculateTotal() {
	for _, product := range o.Products {
		o.total += product.Price * float64(product.Quantity)
	}
}

func (o *Order) UpdateStatus(newStatus string) error {
	validStatuses := []string{"pending", "shipped", "cancelled", "completed"}
	for _, status := range validStatuses {
		if newStatus == status {
			o.Status = newStatus
			return nil
		}
	}
	return errors.New("invalid status")
}
