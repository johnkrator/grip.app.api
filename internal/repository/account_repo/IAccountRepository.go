package account_repo

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"grip.app.api/internal/models"
)

type IAccountRepository interface {
	CreateAccountTx(tx *gorm.DB, account *models.Account) error
	GetByID(id uuid.UUID) (*models.Account, error)
	UpdateBalance(id uuid.UUID, newBalance decimal.Decimal) error
}
