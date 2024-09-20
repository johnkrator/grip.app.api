package email_send_config

import (
	"crypto/tls"
	"fmt"
	"github.com/shopspring/decimal"
	"gopkg.in/mail.v2"
	"os"
	"strconv"
	"time"
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
	fullName := fmt.Sprintf("%s %s", firstName, lastName)
	currentTime := time.Now()
	dateStr := currentTime.Format("Monday, January 02 2006")
	timeStr := currentTime.Format("15:04:05")

	// Function to get currency symbol
	getCurrencySymbol := func(currency string) string {
		switch currency {
		case "NGN":
			return "₦"
		case "USD":
			return "$"
		case "EUR":
			return "€"
		case "GBP":
			return "£"
		default:
			return currency
		}
	}

	currencySymbol := getCurrencySymbol(currency)

	switch transactionType {
	case Registration:
		subject = "Welcome to Grip Finance!"
		body = fmt.Sprintf(`Dear %s,

		Your registration with Grip Finance has been confirmed.
		
		Account Details:
		Account Number: %s
		
		For further enquiries, please contact our customer support through the following channels:
		Email: customerservice@gripfinance.com
		Phone: 07001234567
		
		Thank you for choosing Grip Finance.`, fullName, accountNumber)

	case Deposit:
		subject = "Deposit Confirmation"
		body = fmt.Sprintf(`Dear %s,

		Your deposit of %s%s has been confirmed. Your account will be credited within 5 minutes.
		
		Deposit Details:
		Amount: %s%s
		Date: %s
		Time: %s
		
		For further enquiries, please contact our customer support through the following channels:
		Email: customerservice@gripfinance.com
		Phone: 07001234567
		
		Thank you for choosing Grip Finance.`, fullName, currencySymbol, amount.String(), currencySymbol, amount.String(), dateStr, timeStr)

	case Withdrawal:
		subject = "Withdrawal Confirmation"
		body = fmt.Sprintf(`Dear %s,

		Your withdrawal of %s%s has been confirmed.
		
		Withdrawal Details:
		Amount: %s%s
		Date: %s
		Time: %s
		
		For further enquiries, please contact our customer support through the following channels:
		Email: customerservice@gripfinance.com
		Phone: 07001234567
		
		Thank you for choosing Grip Finance.`, fullName, currencySymbol, amount.Abs().String(), currencySymbol, amount.Abs().String(), dateStr, timeStr)

	case Transfer:
		subject = "Transfer Confirmation"
		body = fmt.Sprintf(`Dear %s,

		Your transfer of %s%s has been confirmed and the recipient is expected to be credited within 5 minutes.
		
		Transfer Details:
		Amount: %s%s
		Date: %s
		Time: %s
		
		For further enquiries, please contact our customer support through the following channels:
		Email: customerservice@gripfinance.com
		Phone: 07001234567
		
		Thank you for choosing Grip Finance.`, fullName, currencySymbol, amount.Abs().String(), currencySymbol, amount.Abs().String(), dateStr, timeStr)

	case FeeCharged:
		subject = "Fee Charged Notification"
		body = fmt.Sprintf(`Dear %s,

		A fee of %s%s has been charged to your account.
		
		Fee Details:
		Amount: %s%s
		Date: %s
		Time: %s
		Description: %s
		
		For further enquiries, please contact our customer support through the following channels:
		Email: customerservice@gripfinance.com
		Phone: 07001234567
		
		Thank you for choosing Grip Finance.`, fullName, currencySymbol, amount.Abs().String(), currencySymbol, amount.Abs().String(), dateStr, timeStr, description)

	default:
		subject = "Grip Finance Transaction Notification"
		body = fmt.Sprintf(`Dear %s,

		A transaction of %s%s has occurred on your account.
		
		Transaction Details:
		Type: %s
		Amount: %s%s
		Date: %s
		Time: %s
		Description: %s
		
		For further enquiries, please contact our customer support through the following channels:
		Email: customerservice@gripfinance.com
		Phone: 07001234567
		
		Thank you for choosing Grip Finance.`, fullName, currencySymbol, amount.Abs().String(), transactionType, currencySymbol, amount.Abs().String(), dateStr, timeStr, description)
	}

	return subject, body
}
