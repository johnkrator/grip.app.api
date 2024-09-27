package config

import (
	"grip.app.api/internal/models"
)

// GetModelsToKeep returns a slice of models to be kept and migrated
func GetModelsToKeep() []interface{} {
	return []interface{}{
		&models.User{},
		&models.UserProfile{},
		&models.Account{},
		&models.Transaction{},
		&models.AuditLog{},
		&models.Card{},
		&models.Investment{},
		&models.Loan{},
		&models.LoanPayment{},
		&models.Notification{},
		&models.Stock{},
		&models.SupportTicket{},
	}
}
