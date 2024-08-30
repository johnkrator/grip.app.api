package models

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/google/uuid"
	"grip.app.api/internal/models/base"
)

type Loan struct {
	base.BaseModel
	UserID       uuid.UUID
	Amount       decimal.Decimal
	InterestRate decimal.Decimal
	Term         string
	Status       string
	Purpose      string
	ApprovedAt   time.Time
}
