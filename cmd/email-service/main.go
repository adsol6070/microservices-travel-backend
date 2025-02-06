package main

import (
	"log"
	"microservices-travel-backend/internal/email-service/adapters"
	"microservices-travel-backend/internal/email-service/services"
	"microservices-travel-backend/pkg/email"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	// Retrieve SendGrid API Key from environment variable
	sendGridAPIKey := os.Getenv("SENDGRID_API_KEY")
	RABBITMQURL := os.Getenv("RABBITMQ_URL")
	if sendGridAPIKey == "" {
		log.Fatalf("SENDGRID_API_KEY is not set")
	}

	// Create EmailClient using SendGrid API Key
	emailClient := email.NewEmailClient(sendGridAPIKey)

	// Create EmailService using the EmailClient
	emailService := services.NewEmailService(emailClient)

	// Set up RabbitMQ connection
	conn, err := amqp.Dial(RABBITMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer channel.Close()

	// Set up RabbitMQ Consumer
	rabbitMQConsumer := adapters.NewRabbitMQConsumer(channel, emailService)

	// Start consuming emails
	rabbitMQConsumer.ConsumeEmails()
}
