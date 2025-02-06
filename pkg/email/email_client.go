package email

import (
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailClient struct {
	APIKey string
}

func NewEmailClient(apiKey string) *EmailClient {
	return &EmailClient{
		APIKey: apiKey,
	}
}

func (c *EmailClient) SendEmail(to, subject, body string) error {
	// Create a new SendGrid client
	client := sendgrid.NewSendClient(c.APIKey)

	// Create a new email message using SendGrid helper methods
	from := mail.NewEmail("Testing", "jatingmttf@gmail.com")
	toEmail := mail.NewEmail("", to)
	content := mail.NewContent("text/plain", body)

	message := mail.NewV3MailInit(from, subject, toEmail, content)

	// Send the email via the SendGrid API
	response, err := client.Send(message)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		return err
	}

	// Log the response to confirm the status
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		fmt.Println("Email sent successfully to:", to)
	} else {
		fmt.Printf("Failed to send email. Status: %v\nResponse Body: %v\n", response.StatusCode, response.Body)
	}

	return nil
}
