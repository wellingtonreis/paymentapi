package repositories

import (
	"context"
	config "wallet/internal/configs"
	domain "wallet/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
)

type WalletRepository struct {
	db *config.MongoDB
}

func NewWalletRepository(db *config.MongoDB) *WalletRepository {
	return &WalletRepository{
		db: db,
	}
}

func (r *WalletRepository) Update(ctx context.Context, wallet *domain.Wallet) error {
	collection := r.db.Client.Database("admin").Collection("payment_order_details")
	filter := bson.M{"user_id": wallet.UserID, "order_id": wallet.OrderID}

	update := bson.M{
		"$set": bson.M{"status": "completed"},
		"$push": bson.M{
			"wallet": bson.M{"balance": wallet.Amount},
		},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
