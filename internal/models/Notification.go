package models

import (
	"github.com/google/uuid"
	"grip.app.api/internal/models/base"
)

type Notification struct {
	base.BaseModel
	UserID  uuid.UUID
	User    User
	Type    string
	Message string
	Read    bool
}
