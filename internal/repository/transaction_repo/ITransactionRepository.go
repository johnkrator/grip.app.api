package transaction_repo

import (
	"github.com/google/uuid"
	"grip.app.api/internal/models"
)

type ITransactionRepository interface {
	Create(transaction *models.Transaction) error
	GetByID(id uuid.UUID) (*models.Transaction, error)
	GetByUserID(userID uuid.UUID) ([]*models.Transaction, error)
	GetByAccountID(accountID uuid.UUID) ([]*models.Transaction, error)
}
