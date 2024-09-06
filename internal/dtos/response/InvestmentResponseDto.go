package response

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type InvestmentResponseDto struct {
	ID           uuid.UUID       `json:"id"`
	Type         string          `json:"type"`
	Amount       decimal.Decimal `json:"amount"`
	CurrentValue decimal.Decimal `json:"currentValue"`
	Status       string          `json:"status"`
}
