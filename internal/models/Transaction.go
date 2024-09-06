package models

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/google/uuid"
	"grip.app.api/internal/models/base"
)

type TransactionType string

const (
	Deposit         TransactionType = "DEPOSIT"
	Withdrawal      TransactionType = "WITHDRAWAL"
	Transfer        TransactionType = "TRANSFER"
	PaymentSent     TransactionType = "PAYMENT_SENT"
	PaymentReceived TransactionType = "PAYMENT_RECEIVED"
	FeeCharged      TransactionType = "FEE_CHARGED"
)

type Transaction struct {
	base.BaseModel
	AccountID   uuid.UUID
	UserID      uuid.UUID
	Type        TransactionType
	Amount      decimal.Decimal
	Currency    string
	Description string
	Category    string
	Status      string
	Timestamp   time.Time
}
