package email_send_config

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/mail.v2"
	"os"
	"strconv"
)

func SendPasswordResetEmail(toEmail, firstName, resetLink string) error {
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
	m.SetHeader("Subject", "Password Reset Request - Grip Finance")

	body := fmt.Sprintf(`
		Dear %s,

		We received a request to reset your password for your Grip Finance account.

		To reset your password, please click on the following link:
		%s

		This link will expire in 1 hour for security reasons.

		If you didn't request a password reset, please ignore this email or contact our support team if you have any concerns.

		Best regards,
		Team Grip Finance
	`, firstName, resetLink)

	m.SetBody("text/plain", body)

	d := mail.NewDialer(smtpServer, smtpPort, smtpUser, smtpPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	return nil
}
