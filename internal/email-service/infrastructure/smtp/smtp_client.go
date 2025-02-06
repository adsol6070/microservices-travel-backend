package smtp

import (
	"fmt"
	"net/smtp"
	"os"
)

// SMTPClient handles sending emails
type SMTPClient struct {
	SMTPServer string
	Port       string
	Username   string
	Password   string
}

// NewSMTPClient initializes an SMTP client
func NewSMTPClient() *SMTPClient {
	return &SMTPClient{
		SMTPServer: os.Getenv("SMTP_SERVER"),
		Port:       os.Getenv("SMTP_PORT"),
		Username:   os.Getenv("SMTP_USERNAME"),
		Password:   os.Getenv("SMTP_PASSWORD"),
	}
}

// SendEmail sends an email using the SMTP client
func (s *SMTPClient) SendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.Username, s.Password, s.SMTPServer)
	msg := []byte("Subject: " + subject + "\r\n" + "\r\n" + body + "\r\n")

	err := smtp.SendMail(s.SMTPServer+":"+s.Port, auth, s.Username, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	fmt.Println("Email sent successfully to", to)
	return nil
}
