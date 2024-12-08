package domain

type Wallet struct {
	UserID  string  `json:"user_id" validate:"required"`
	OrderID string  `json:"order_id" validate:"required"`
	Amount  float64 `json:"amount" validate:"required"`
}
