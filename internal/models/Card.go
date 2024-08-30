package models

import (
	"time"

	"github.com/google/uuid"
	"grip.app.api/internal/models/base"
)

type Card struct {
	base.BaseModel
	AccountID      uuid.UUID
	CardNumber     string
	CardType       string
	ExpirationDate time.Time
	CVV            string
	Status         string
}
