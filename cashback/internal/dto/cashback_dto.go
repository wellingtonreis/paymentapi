package dto

type Notification struct {
	OrderID string  `json:"order_id" validate:"required"`
	Amount  float64 `json:"amount" validate:"required"`
	Status  string  `json:"status" validate:"required"`
}
