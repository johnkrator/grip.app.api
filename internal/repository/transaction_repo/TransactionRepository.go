package transaction_repo

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"grip.app.api/internal/models"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *TransactionRepository) GetByID(id uuid.UUID) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) GetByUserID(userID uuid.UUID) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.Where("user_id = ?", userID).Find(&transactions).Error
	return transactions, err
}

func (r *TransactionRepository) GetByAccountID(accountID uuid.UUID) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.Where("account_id = ?", accountID).Find(&transactions).Error
	return transactions, err
}
