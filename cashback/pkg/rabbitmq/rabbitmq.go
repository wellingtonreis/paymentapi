package rabbitmq

import (
	"fmt"

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

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery, exchange string, exchangeType string, queueName string, key string) error {

	if err := ch.ExchangeDeclare(
		exchange,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		fmt.Println("Erro ao declarar o exchange")
		return err
	}

	queue, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("Erro ao declarar a fila")
		return err
	}

	if err = ch.QueueBind(
		queue.Name,
		key,
		exchange,
		false,
		nil,
	); err != nil {
		fmt.Println("Erro ao fazer o bind")
		return err
	}

	msgs, err := ch.Consume(
		queue.Name,
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("Erro ao consumir a fila")
		return err
	}
	for msg := range msgs {
		out <- msg
	}
	return nil
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
