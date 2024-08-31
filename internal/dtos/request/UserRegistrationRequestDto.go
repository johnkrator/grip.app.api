package request

import (
	"time"
)

type UserRegistrationRequestDto struct {
	FirstName     string    `json:"firstName" binding:"required"`
	LastName      string    `json:"lastName" binding:"required"`
	Email         string    `json:"email" binding:"required,email"`
	PhoneNumber   string    `json:"phoneNumber" binding:"required"`
	DateOfBirth   time.Time `json:"dateOfBirth" binding:"required"`
	Address       string    `json:"address" binding:"required"`
	Password      string    `json:"password" binding:"required,min=8"`
	Occupation    string    `json:"occupation" binding:"required"`
	IncomeRange   string    `json:"incomeRange" binding:"required"`
	RiskTolerance string    `json:"riskTolerance" binding:"required"`
	Preferences   string    `json:"preferences"`
}
