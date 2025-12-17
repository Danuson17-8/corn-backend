package utils

import (
	"net/smtp"
	"os"

	"github.com/Danuson17-8/corn-backend/config"
)

func SendEmail(cfg *config.EnvConfig, to string, subject string, body string) error {
	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASS")

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// Gmail SMTP
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from,
		[]string{to},
		[]byte(msg),
	)

	return err
}
