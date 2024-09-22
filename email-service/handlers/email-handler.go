package handlers

import (
    "encoding/json"
    "log"
    "email-service/service"
    "email-service/model"
    "github.com/streadway/amqp"
)

func HandleEmailQueue(ch *amqp.Channel, q amqp.Queue, emailService *service.EmailService) {
    msgs, err := ch.Consume(
        q.Name, "", true, false, false, false, nil,
    )
    if err != nil {
        log.Fatalf("Failed to consume from queue: %v", err)
    }

    forever := make(chan bool)

    go func() {
        for d := range msgs {
            var payload model.EmailPayload
            if err := json.Unmarshal(d.Body, &payload); err != nil {
                log.Printf("Error decoding JSON: %v", err)
                continue
            }

            // Send email
            if err := emailService.SendEmail(payload); err != nil {
                log.Printf("Failed to send email: %v", err)
            } else {
                log.Printf("Email sent to %s", payload.To)
            }
        }
    }()

    log.Println("Waiting for messages...")
    <-forever
}
