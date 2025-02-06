package email

import (
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
	done := make(chan struct{})
	go s.publisher.Publish(done, message, "emailQueue", os.Getenv("RABBITMQ_URL"))

	<-done
	return nil
}
