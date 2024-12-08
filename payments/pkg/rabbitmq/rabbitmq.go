package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel(url string) (*amqp.Channel, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}

func Publish(ch *amqp.Channel, body string, exName string, key string) error {
	err := ch.Publish(
		exName,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
