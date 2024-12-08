package configs

import (
	"context"
	"fmt"
	"time"

	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"

	viper "github.com/spf13/viper"
)

type MongoDB struct {
	Client *mongo.Client
}

func uri() string {
	driver := viper.GetString("DB_DRIVER")
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")

	mongodbUri := fmt.Sprintf("%s://%s:%s", driver, host, port)
	return mongodbUri
}

func ConnectDB() (*MongoDB, error) {

	mongodbUri := uri()

	clientOptions := options.Client().ApplyURI(mongodbUri)
	clientOptions.SetConnectTimeout(10 * time.Second)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return &MongoDB{Client: client}, nil
}
