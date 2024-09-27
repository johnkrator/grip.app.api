package response

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type StockResponse struct {
	ID           uuid.UUID       `json:"id"`
	Symbol       string          `json:"symbol"`
	Name         string          `json:"name"`
	CurrentPrice decimal.Decimal `json:"currentPrice"`
	LastUpdated  time.Time       `json:"lastUpdated"`
}
