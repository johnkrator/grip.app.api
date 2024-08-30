package models

import (
	"github.com/google/uuid"
	"grip.app.api/internal/models/base"
)

type UserProfile struct {
	base.BaseModel
	UserID        uuid.UUID `gorm:"uniqueIndex"`
	Occupation    string
	IncomeRange   string
	RiskTolerance string
	Preferences   string
}
