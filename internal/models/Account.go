package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"grip.app.api/internal/models/base"
)

type AccountType string
type Currency string
type AccountStatus string

const (
	CurrentAccount AccountType = "current"
	SavingsAccount AccountType = "savings"
)

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
)

const (
	Active   AccountStatus = "active"
	Inactive AccountStatus = "inactive"
)

type Account struct {
	base.BaseModel
	UserID        uuid.UUID
	AccountType   AccountType
	AccountNumber string
	Balance       decimal.Decimal
	Currency      Currency
	Status        AccountStatus
	InterestRate  decimal.Decimal
}
