package rabbitmq

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type RabbitMQClient struct {
	conn *amqp.Connection
}

func NewRabbitMQClient(addr string) (*RabbitMQClient, error) {
	conn, err := amqp.Dial(addr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &RabbitMQClient{
		conn: conn,
	}, nil
}

func (r *RabbitMQClient) Push(name string, data interface{}) error {
	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = ch.PublishWithContext(ctx, "", queue.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	if err != nil {
		return err
	}
	return nil
}
