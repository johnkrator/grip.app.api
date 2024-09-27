package request

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type CreateInvestmentRequest struct {
	UserID       uuid.UUID       `json:"userId"`
	Type         string          `json:"type"`
	Amount       decimal.Decimal `json:"amount"`
	PurchaseDate time.Time       `json:"purchaseDate"`
}
