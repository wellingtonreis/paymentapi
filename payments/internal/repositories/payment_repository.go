package repositories

import (
	"context"
	"errors"
	config "payments/internal/configs"
	domain "payments/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
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

	return insertedID.Hex(), nil
}

func (r *PaymentRepository) Update(ctx context.Context, payment *domain.Order) error {
	collection := r.db.Client.Database("admin").Collection("payment_order_details")
	filter := bson.M{"user_id": payment.UserID, "order_id": payment.OrderID}

	update := bson.M{
		"$set": bson.M{"status": payment.Status},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
