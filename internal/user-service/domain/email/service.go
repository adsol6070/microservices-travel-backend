package email

import (
	"log"
	"microservices-travel-backend/internal/shared/rabbitmq"
	"os"
)

type EmailService struct {
	publisher *rabbitmq.Client
}

func NewEmailService(publisher *rabbitmq.Client) *EmailService {
	return &EmailService{
		publisher: publisher,
	}
}

func (s *EmailService) SendEmail(message []byte) error {
	go func() {
		if err := s.publisher.Publish(message, "emailQueue", os.Getenv("RABBITMQ_URL")); err != nil {
			log.Printf("Error while publishing email message: %v", err)
		}

	}()

	return nil
}
