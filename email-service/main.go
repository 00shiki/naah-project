package main

import (
    "log"
    "email-service/service"
    "email-service/handlers"
    "github.com/streadway/amqp"
	"os"
)

func main() {
    // Load Email Service
    emailService, err := service.NewEmailService()
    if err != nil {
        log.Fatalf("Error initializing email service: %v", err)
    }

    // Connect to RabbitMQ
    conn, err := amqp.Dial(os.Getenv("RABBIT_MQ_URL"))
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open RabbitMQ channel: %v", err)
    }
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "email_queue", false, false, false, false, nil,
    )
    if err != nil {
        log.Fatalf("Failed to declare RabbitMQ queue: %v", err)
    }

    // Handle messages from RabbitMQ
    handlers.HandleEmailQueue(ch, q, emailService)
}
