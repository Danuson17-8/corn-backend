package utils

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/Danuson17-8/corn-backend/config"
)

func SendEmail(cfg *config.EnvConfig, to, subject, body string) error {
	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASS")
	auth := smtp.PlainAuth(
		"",
		from,
		password,
		"smtp.gmail.com",
	)

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		from, to, subject, body)

	return smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(msg))
}
