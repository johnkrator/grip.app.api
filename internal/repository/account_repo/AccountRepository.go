package account_repo

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"grip.app.api/internal/models"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) IAccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) CreateAccountTx(tx *gorm.DB, account *models.Account) error {
	return tx.Create(account).Error
}

func (r *AccountRepository) GetByID(id uuid.UUID) (*models.Account, error) {
	var account models.Account
	err := r.db.First(&account, id).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) UpdateBalance(id uuid.UUID, newBalance decimal.Decimal) error {
	return r.db.Model(&models.Account{}).Where("id = ?", id).Update("balance", newBalance).Error
}
