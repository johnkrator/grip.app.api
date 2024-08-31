package account_repo

import (
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
