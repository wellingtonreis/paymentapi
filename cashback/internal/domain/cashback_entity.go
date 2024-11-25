package domain

import (
	"errors"
)

type Cashback struct {
	Amount          float64 `bson:"amount"`
	DiscountPercent float64 `bson:"discount_percent"`
}

func NewCashback(amount, discountPercent float64) (*Cashback, error) {
	if amount < 0 {
		return nil, errors.New("amount must be greater than or equal to zero")
	}

	if discountPercent < 0 {
		return nil, errors.New("discountPercent must be greater than or equal to zero")
	}

	return &Cashback{
		Amount:          amount,
		DiscountPercent: discountPercent,
	}, nil
}

func (c *Cashback) Calculate() float64 {
	return c.Amount * c.DiscountPercent
}
