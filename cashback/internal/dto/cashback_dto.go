package dto

type Notification struct {
	UserID  string  `json:"user_id" validate:"required"`
	OrderID string  `json:"order_id" validate:"required"`
	Amount  float64 `json:"amount" validate:"required"`
	Status  string  `json:"status" validate:"required"`
}
