package models

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/google/uuid"
	"grip.app.api/internal/models/base"
)

type LoanPayment struct {
	base.BaseModel
	LoanID  uuid.UUID
	Amount  decimal.Decimal
	DueDate time.Time
	PaidAt  time.Time
	Status  string
}
