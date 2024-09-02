package response

import (
	"github.com/google/uuid"
	"time"
)

type ChangePasswordResponse struct {
	ID           uuid.UUID
	FirstName    string
	LastName     string
	Email        string `gorm:"uniqueIndex"`
	PhoneNumber  string
	DateOfBirth  time.Time
	Address      string
	AccessToken  string `gorm:"default:null"`
	RefreshToken string `gorm:"default:null"`
	IsVerified   bool   `gorm:"default:false"`
	IsAdmin      bool   `gorm:"default:false"`
	IsDeleted    bool   `gorm:"default:false"`
	Role         Role   `gorm:"default:customer"`
	Message      string `json:"message"`
}
