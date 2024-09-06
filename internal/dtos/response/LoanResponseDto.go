package response

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type LoanResponseDto struct {
	ID           uuid.UUID       `json:"id"`
	Amount       decimal.Decimal `json:"amount"`
	InterestRate decimal.Decimal `json:"interestRate"`
	Term         string          `json:"term"`
	Status       string          `json:"status"`
	Purpose      string          `json:"purpose"`
}
