// bagian yg dikomen di bawah ini ditambahin di main.go nya

// rabbitMQ, err := rabbitmq.NewRabbitMQClient(os.Getenv("RABBITMQ_ADDR"))
// if err != nil {
// 	log.Fatalf("failed to create rabbitmq client: %v", err)
// }

package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
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
	log.Println("SUCCESS INITIATE RABBITMQ CLIENT")
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

// inputnya
// - nama
// - email
// - list of
//   - shoe name
//   - shoe size
//   - shoe qty
// -  total price paid

type OrderDetail struct {
	ShoeName string `json:"shoe_name"`
	ShoeSize int64  `json:"shoe_size"`
	Qty      int32  `json:"quantity"`
}

type OrderReceiptEmail struct {
	UserName    string        `json:"user_name"`
	UserEmail   string        `json:"user_email"`
	OrderId     int64         `json:"order_id"`
	TotalPrice  int32         `json:"total_price"`
	OrderDetail []OrderDetail `json:"order_detail"`
}

func SendReceiptEmail(order OrderReceiptEmail, rabbitMQ *RabbitMQClient) error {
	emailPayload := map[string]interface{}{
		"to":            order.UserEmail,
		"subject":       "Your Order Receipt",
		"type":          "receipt",
		"order_receipt": order,
	}

	err := rabbitMQ.Push("email_queue", emailPayload)
	if err != nil {
		log.Printf("Error sending email receipt: %v", err)
		return fmt.Errorf("cannot push email")
	}
	return nil
}

func SendDeliveredEmail(userEmail string, orderID int, rabbitMQ *RabbitMQClient) error {
	emailPayload := map[string]interface{}{
		"to":       userEmail,
		"subject":  "Your Order Has Been Delivered",
		"type":     "delivered",
		"order_id": orderID,
	}

	err := rabbitMQ.Push("email_queue", emailPayload)
	if err != nil {
		log.Printf("Error sending email receipt: %v", err)
		return fmt.Errorf("cannot push email")
	}
	return nil
}
