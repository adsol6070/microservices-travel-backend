package adapters

import (
    "encoding/json"
    "log"
    "microservices-travel-backend/internal/email-service/services"
    "github.com/streadway/amqp"
)

type RabbitMQConsumer struct {
    channel *amqp.Channel
    emailService *services.EmailService
}

func NewRabbitMQConsumer(channel *amqp.Channel, emailService *services.EmailService) *RabbitMQConsumer {
    return &RabbitMQConsumer{
        channel: channel,
        emailService: emailService,
    }
}

func (r *RabbitMQConsumer) ConsumeEmails() {
    // Declare the queue to consume from
    queue, err := r.channel.QueueDeclare(
        "email_queue", // Queue name
        true,          // Durable
        false,         // Auto-delete
        false,         // Exclusive
        false,         // No-wait
        nil,           // Arguments
    )
    if err != nil {
        log.Fatalf("Failed to declare a queue: %v", err)
    }

    // Start consuming messages from the queue
    msgs, err := r.channel.Consume(
        queue.Name, // Queue name
        "",         // Consumer tag
        true,        // Auto-ack
        false,       // Exclusive
        false,       // No-local
        false,       // No-wait
        nil,         // Arguments
    )
    if err != nil {
        log.Fatalf("Failed to register a consumer: %v", err)
    }

    log.Println("Waiting for messages. To exit press CTRL+C")

    // Start listening for messages
    for msg := range msgs {
        var emailMessage struct {
            To      string `json:"to"`
            Subject string `json:"subject"`
            Body    string `json:"body"`
        }
        err := json.Unmarshal(msg.Body, &emailMessage)
        if err != nil {
            log.Printf("Failed to unmarshal message: %v", err)
            continue
        }

        err = r.emailService.SendCustomEmail(emailMessage.To, emailMessage.Subject, emailMessage.Body)
        if err != nil {
            log.Printf("Failed to send email: %v", err)
        }
    }
}
