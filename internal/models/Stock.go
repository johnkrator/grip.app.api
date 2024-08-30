package models

import (
	"github.com/shopspring/decimal"
	"time"

	"grip.app.api/internal/models/base"
)

type Stock struct {
	base.BaseModel
	Symbol       string
	Name         string
	CurrentPrice decimal.Decimal
	LastUpdated  time.Time
}
