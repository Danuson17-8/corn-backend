package utils

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(to, subject, body string) error {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	email := os.Getenv("EMAIL_USER")
	from := mail.NewEmail("Corn Cornn", email)
	toEmail := mail.NewEmail("", to)
	message := mail.NewSingleEmail(from, subject, toEmail, body, body)
	client := sendgrid.NewSendClient(apiKey)
	resp, err := client.Send(message)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("SendGrid error %d: %s", resp.StatusCode, resp.Body)
	}
	return nil
}
