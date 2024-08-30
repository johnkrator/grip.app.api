package models

import (
	"github.com/google/uuid"
	"grip.app.api/internal/models/base"
	"time"
)

type AuditLog struct {
	base.BaseModel
	UserID     uuid.UUID
	User       User
	Action     string
	EntityType string
	EntityID   uuid.UUID
	Details    string
	Timestamp  time.Time
}
