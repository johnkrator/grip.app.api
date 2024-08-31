package account_repo

import (
	"gorm.io/gorm"
	"grip.app.api/internal/models"
)

type IAccountRepository interface {
	CreateAccountTx(tx *gorm.DB, account *models.Account) error
}
