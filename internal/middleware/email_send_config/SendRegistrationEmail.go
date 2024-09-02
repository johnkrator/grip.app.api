package email_send_config

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/mail.v2"
	"os"
	"strconv"
)

func SendRegistrationEmail(toEmail, firstName, lastName, accountNumber, token string) error {
	fromEmail := os.Getenv("FROM_EMAIL")
	if fromEmail == "" {
		return fmt.Errorf("FROM_EMAIL environment variable is not set")
	}

	smtpServer := os.Getenv("SMTP_SERVER")
	if smtpServer == "" {
		return fmt.Errorf("SMTP_SERVER environment variable is not set")
	}

	smtpPortStr := os.Getenv("SMTP_PORT")
	if smtpPortStr == "" {
		return fmt.Errorf("SMTP_PORT environment variable is not set")
	}
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		return fmt.Errorf("invalid SMTP_PORT: %v", err)
	}

	smtpUser := os.Getenv("SMTP_USER")
	if smtpUser == "" {
		return fmt.Errorf("SMTP_USER environment variable is not set")
	}

	smtpPassword := os.Getenv("SMTP_PASSWORD")
	if smtpPassword == "" {
		return fmt.Errorf("SMTP_PASSWORD environment variable is not set")
	}

	m := mail.NewMessage()
	m.SetHeader("From", fromEmail)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Welcome to Grip Finance!")

	body := fmt.Sprintf(`
		Dear %s %s,

		Welcome to Grip Finance! We're thrilled to have you on board.

		Your registration was successful, and your account is now active. Here are your account details:

		Account Number: %s
		Login Token: %s

		Please use this token for your first login. For security reasons, a new token will be generated after each successful login.

		If you have any questions or need assistance, please don't hesitate to contact our support team.

		Best regards,
		Team Grip Finance
	`, firstName, lastName, accountNumber, token)

	m.SetBody("text/plain", body)

	d := mail.NewDialer(smtpServer, smtpPort, smtpUser, smtpPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	return nil
}
