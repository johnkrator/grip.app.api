package response

import (
	"github.com/google/uuid"
	"time"
)

// Role type enum
type Role string

const (
	AdminRole    Role = "admin"
	ManagerRole  Role = "manager"
	EmployeeRole Role = "employee"
	CustomerRole Role = "customer"
)

type UserRegistrationResponseDto struct {
	ID          uuid.UUID              `json:"id"`
	FirstName   string                 `json:"firstName"`
	LastName    string                 `json:"lastName"`
	Email       string                 `json:"email"`
	PhoneNumber string                 `json:"phoneNumber"`
	DateOfBirth time.Time              `json:"dateOfBirth"`
	Address     string                 `json:"address"`
	Role        Role                   `json:"role"`
	Profile     UserProfileResponseDto `json:"profile"`
	Account     AccountResponseDto     `json:"account"`
}

type UserProfileResponseDto struct {
	Occupation    string `json:"occupation"`
	IncomeRange   string `json:"incomeRange"`
	RiskTolerance string `json:"riskTolerance"`
	Preferences   string `json:"preferences"`
}
