package response

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PortfolioResponse struct {
	UserID              uuid.UUID          `json:"userId"`
	TotalPortfolioValue decimal.Decimal    `json:"totalPortfolioValue"`
	Investments         []InvestmentDetail `json:"investments"`
}

type InvestmentDetail struct {
	ID         uuid.UUID       `json:"id"`
	Type       string          `json:"type"`
	TotalValue decimal.Decimal `json:"totalValue"`
	Holdings   []StockHolding  `json:"holdings,omitempty"`
}

type StockHolding struct {
	StockID      string          `json:"stockId"`
	CompanyName  string          `json:"companyName"`
	Shares       decimal.Decimal `json:"shares"`
	CurrentPrice decimal.Decimal `json:"currentPrice"`
	TotalValue   decimal.Decimal `json:"totalValue"`
}
