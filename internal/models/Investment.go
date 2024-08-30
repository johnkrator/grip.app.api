package models

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/google/uuid"
	"grip.app.api/internal/models/base"
)

type Investment struct {
	base.BaseModel
	UserID       uuid.UUID
	Type         string
	Amount       decimal.Decimal
	PurchaseDate time.Time
	CurrentValue decimal.Decimal
	Status       string
}
