package models

import (
	"github.com/google/uuid"
	"grip.app.api/internal/models/base"
)

type SupportTicket struct {
	base.BaseModel
	UserID      uuid.UUID
	Subject     string
	Description string
	Status      string
	Priority    string
}
