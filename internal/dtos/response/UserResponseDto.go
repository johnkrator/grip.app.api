package response

import (
	"github.com/google/uuid"
	"time"
)

type UserResponseDto struct {
	ID          uuid.UUID               `json:"id"`
	FirstName   string                  `json:"firstName"`
	LastName    string                  `json:"lastName"`
	Email       string                  `json:"email"`
	PhoneNumber string                  `json:"phoneNumber"`
	DateOfBirth time.Time               `json:"dateOfBirth"`
	Address     string                  `json:"address"`
	Role        Role                    `json:"role"`
	IsVerified  bool                    `json:"isVerified"`
	Profile     *UserProfileResponseDto `json:"profile,omitempty"`
}
