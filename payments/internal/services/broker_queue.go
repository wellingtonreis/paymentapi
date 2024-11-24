package services

import (
	"fmt"
	rabbitmq "payments/pkg/rabbitmq"

	viper "github.com/spf13/viper"
)

func SendMessage(notification string) (bool, error) {

	url := viper.GetString("AMQP.URL")
	exchenge := viper.GetString("AMQP.EXCHANGE")
	queue := viper.GetString("AMQP.QUEUE")

	ch, err := rabbitmq.OpenChannel(url)
	if err != nil {
		return false, fmt.Errorf("falha ao abrir o canal com rabbitmq: %v", err)
	}

	err = rabbitmq.Publish(ch, notification, exchenge, queue)
	if err != nil {
		return false, fmt.Errorf("falha ao publicar o JSON no RabbitMQ %v", err)
	}

	return true, nil
}
