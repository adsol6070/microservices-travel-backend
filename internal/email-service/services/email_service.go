package services

import (
	"encoding/json"
	"log"
	"microservices-travel-backend/internal/email-service/domain/models"
	"microservices-travel-backend/internal/email-service/infrastructure/sendgrid"

	amqp "github.com/rabbitmq/amqp091-go"
)

// EmailService processes email messages
type EmailService struct {
	SendGridClient *sendgrid.SendGridClient
}

// NewEmailService initializes an email service
func NewEmailService(sendGridClient *sendgrid.SendGridClient) *EmailService {
	return &EmailService{SendGridClient: sendGridClient}
}

// ProcessEmailMessage processes a RabbitMQ message and sends an email
func (e *EmailService) ProcessEmailMessage(msg amqp.Delivery) {
	var email models.Email
	err := json.Unmarshal(msg.Body, &email)
	if err != nil {
		log.Println("Failed to parse email message:", err)
		return
	}

	err = e.SendGridClient.SendEmail(email.To, email.Subject, email.Body)
	if err != nil {
		log.Println("Failed to send email:", err)
	} else {
		log.Println("Email successfully sent to", email.To)
	}
}
