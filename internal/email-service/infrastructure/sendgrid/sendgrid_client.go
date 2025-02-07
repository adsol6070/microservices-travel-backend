package sendgrid

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendGridClient handles sending emails via SendGrid
type SendGridClient struct {
	APIKey string
	Sender string
}

// NewSendGridClient initializes a SendGrid client
func NewSendGridClient() *SendGridClient {
	return &SendGridClient{
		APIKey: os.Getenv("SENDGRID_API_KEY"),
		Sender: os.Getenv("SENDER_EMAIL"),
	}
}

// SendEmail sends an email using SendGrid
func (s *SendGridClient) SendEmail(to, subject, body string) error {
	from := mail.NewEmail("No Reply", s.Sender)
	toEmail := mail.NewEmail("", to)
	message := mail.NewSingleEmail(from, subject, toEmail, body, body)
	client := sendgrid.NewSendClient(s.APIKey)

	response, err := client.Send(message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Println("Email sent successfully to", to, "Status Code:", response.StatusCode)
	return nil
}
