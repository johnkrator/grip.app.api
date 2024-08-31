package models

import (
	"github.com/google/uuid"
	"grip.app.api/internal/models/base"
)

type IncomeRange string

const (
	IncomeRangeLow    IncomeRange = "low"
	IncomeRangeMedium IncomeRange = "medium"
	IncomeRangeHigh   IncomeRange = "high"
)

type UserProfile struct {
	base.BaseModel
	UserID        uuid.UUID `gorm:"uniqueIndex"`
	Occupation    string
	IncomeRange   IncomeRange
	RiskTolerance string
	Preferences   string
}
