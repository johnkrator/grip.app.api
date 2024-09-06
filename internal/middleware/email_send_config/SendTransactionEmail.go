package email_send_config

import (
	"crypto/tls"
	"fmt"
	"github.com/shopspring/decimal"
	"gopkg.in/mail.v2"
	"os"
	"strconv"
)

type TransactionType string

const (
	Registration TransactionType = "Registration"
	Deposit      TransactionType = "Deposit"
	Withdrawal   TransactionType = "Withdrawal"
	Transfer     TransactionType = "Transfer"
	FeeCharged   TransactionType = "FeeCharged"
)

func SendTransactionEmail(toEmail, firstName, lastName, accountNumber string, transactionType TransactionType, amount decimal.Decimal, currency, description string) error {
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

	subject, body := getEmailContent(firstName, lastName, accountNumber, transactionType, amount, currency, description)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := mail.NewDialer(smtpServer, smtpPort, smtpUser, smtpPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	return nil
}

func getEmailContent(firstName, lastName, accountNumber string, transactionType TransactionType, amount decimal.Decimal, currency, description string) (string, string) {
	var subject, body string

	switch transactionType {
	case Registration:
		subject = "Welcome to Grip Finance!"
		body = fmt.Sprintf(`
			Dear %s %s,
			
			Welcome to Grip Finance! We're thrilled to have you on board.
			
			Your registration was successful, and your account is now active. Here are your account details:
			
			Account Number: %s
			
			If you have any questions or need assistance, please don't hesitate to contact our support team.
			
			Best regards,
			Team Grip Finance
`, firstName, lastName, accountNumber)

	case Deposit:
		subject = "Deposit Confirmation"
		body = fmt.Sprintf(`
			Dear %s %s,
			
			We're pleased to confirm that a deposit has been made to your Grip Finance account.
			
			Account Number: %s
			Amount Deposited: %s %s
			Description: %s
			
			If you have any questions about this transaction, please contact our support team.
			
			Thank you for choosing Grip Finance!
			
			Best regards,
			Team Grip Finance
`, firstName, lastName, accountNumber, amount.String(), currency, description)

	case Withdrawal:
		subject = "Withdrawal Confirmation"
		body = fmt.Sprintf(`
			Dear %s %s,
			
			This email confirms a withdrawal from your Grip Finance account.
			
			Account Number: %s
			Amount Withdrawn: %s %s
			Description: %s
			
			If you did not authorize this withdrawal, please contact our support team immediately.
			
			Thank you for using Grip Finance!
			
			Best regards,
			Team Grip Finance
`, firstName, lastName, accountNumber, amount.Abs().String(), currency, description)

	case Transfer:
		subject = "Transfer Confirmation"
		body = fmt.Sprintf(`
			Dear %s %s,
			
			A transfer has been processed from your Grip Finance account.
			
			Account Number: %s
			Amount Transferred: %s %s
			Description: %s
			
			If you did not authorize this transfer, please contact our support team immediately.
			
			Thank you for using Grip Finance!
			
			Best regards,
			Team Grip Finance
`, firstName, lastName, accountNumber, amount.Abs().String(), currency, description)

	case FeeCharged:
		subject = "Fee Charged Notification"
		body = fmt.Sprintf(`
			Dear %s %s,
			
			A fee has been charged to your Grip Finance account.
			
			Account Number: %s
			Fee Amount: %s %s
			Description: %s
			
			If you have any questions about this fee, please contact our support team.
			
			Thank you for using Grip Finance!
			
			Best regards,
			Team Grip Finance
`, firstName, lastName, accountNumber, amount.Abs().String(), currency, description)

	default:
		subject = "Grip Finance Transaction Notification"
		body = fmt.Sprintf(`
			Dear %s %s,
			
			A transaction has occurred on your Grip Finance account.
			
			Account Number: %s
			Transaction Type: %s
			Amount: %s %s
			Description: %s
			
			If you have any questions about this transaction, please contact our support team.
			
			Thank you for using Grip Finance!
			
			Best regards,
			Team Grip Finance
`, firstName, lastName, accountNumber, transactionType, amount.Abs().String(), currency, description)
	}

	return subject, body
}
