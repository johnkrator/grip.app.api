package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SetupRelations(db *gorm.DB) {
	// User relations
	db.Model(&User{}).Association("Accounts")
	db.Model(&User{}).Association("Transactions")
	db.Model(&User{}).Association("UserProfile")
	db.Model(&User{}).Association("Loans")
	db.Model(&User{}).Association("Investments")
	db.Model(&User{}).Association("SupportTickets")
	db.Model(&User{}).Association("Notifications")
	db.Model(&User{}).Association("AuditLogs")

	// Account relations
	db.Model(&Account{}).Association("Transactions")
	db.Model(&Account{}).Association("Cards")

	// Loan relations
	db.Model(&Loan{}).Association("LoanPayments")

	// Investment and Stock relations
	_ = db.SetupJoinTable(&Investment{}, "Stocks", &InvestmentStock{})
}

// InvestmentStock Define a join table struct for the many-to-many relationship
type InvestmentStock struct {
	InvestmentID uuid.UUID `gorm:"primaryKey"`
	StockID      uuid.UUID `gorm:"primaryKey"`
}
