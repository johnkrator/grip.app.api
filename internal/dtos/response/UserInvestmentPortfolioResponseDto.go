package response

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type UserInvestmentPortfolioResponse struct {
	UserID              uuid.UUID               `json:"userId"`
	TotalPortfolioValue decimal.Decimal         `json:"totalPortfolioValue"`
	Investments         []InvestmentResponseDto `json:"investments"`
}
