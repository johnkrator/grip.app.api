package request

import (
	"time"
)

type UpdateUserRequestDto struct {
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	PhoneNumber   string    `json:"phoneNumber"`
	DateOfBirth   time.Time `json:"dateOfBirth"`
	Address       string    `json:"address"`
	Occupation    string    `json:"occupation"`
	IncomeRange   string    `json:"incomeRange"`
	RiskTolerance string    `json:"riskTolerance"`
	Preferences   string    `json:"preferences"`
}
