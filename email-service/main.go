package main

import (
	"email-service/handlers"
	"email-service/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/streadway/amqp"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		http.HandleFunc("/healthz", healthCheckHandler)
		log.Printf("Starting HTTP server on port %s...\n", port)
		if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// Initialize the email service
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

	// Declare RabbitMQ queue
	q, err := ch.QueueDeclare(
		"email_queue", // Queue name
		false,         // Durable
		false,         // Delete when unused
		false,         // Exclusive
		false,         // No-wait
		nil,           // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare RabbitMQ queue: %v", err)
	}

	// Start handling messages from RabbitMQ queue
	wg.Add(1)
	go func() {
		defer wg.Done()
		handlers.HandleEmailQueue(ch, q, emailService)
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}

// Health check handler
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Service is healthy")
}
