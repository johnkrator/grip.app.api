package request

import (
	"github.com/shopspring/decimal"
)

type UpdateInvestmentRequest struct {
	Type         string          `json:"type"`
	Amount       decimal.Decimal `json:"amount"`
	CurrentValue decimal.Decimal `json:"currentValue"`
	Status       string          `json:"status"`
}
