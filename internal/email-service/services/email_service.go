package services

import (
    "fmt"
    "microservices-travel-backend/pkg/email"
)

type EmailService struct {
    client *email.EmailClient
}

func NewEmailService(client *email.EmailClient) *EmailService {
    return &EmailService{
        client: client,
    }
}

// SendCustomEmail allows sending a custom email (used for other notifications)
func (s *EmailService) SendCustomEmail(to, subject, body string) error {
    err := s.client.SendEmail(to, subject, body)
    if err != nil {
        return fmt.Errorf("failed to send custom email: %v", err)
    }
    return nil
}
