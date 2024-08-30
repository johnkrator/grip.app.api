package models

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/google/uuid"
	"grip.app.api/internal/models/base"
)

type Transaction struct {
	base.BaseModel
	AccountID   uuid.UUID
	UserID      uuid.UUID
	Type        string
	Amount      decimal.Decimal
	Currency    string
	Description string
	Category    string
	Status      string
	Timestamp   time.Time
}
