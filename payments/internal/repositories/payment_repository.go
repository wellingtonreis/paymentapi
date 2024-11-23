package repositories

import (
	"context"
	"errors"
	config "payments/internal/configs"
	domain "payments/internal/domain"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentRepository struct {
	db *config.MongoDB
}

func NewPaymentRepository(db *config.MongoDB) *PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (r *PaymentRepository) Save(ctx context.Context, payment *domain.Order) (string, error) {

	res, err := r.db.Client.Database("admin").Collection("payment_order_details").InsertOne(ctx, payment)
	if err != nil {
		return "", err
	}
	insertedID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("error casting to object id")
	}

	r.db.Client.Disconnect(ctx)
	return insertedID.Hex(), nil
}
