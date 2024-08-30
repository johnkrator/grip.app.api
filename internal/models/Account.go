package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"grip.app.api/internal/models/base"
)

type Account struct {
	base.BaseModel
	UserID        uuid.UUID
	AccountType   string
	AccountNumber string
	Balance       decimal.Decimal
	Currency      string
	Status        string
	InterestRate  decimal.Decimal
}
