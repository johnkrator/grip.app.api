package user_repo

import (
	"gorm.io/gorm"
	"grip.app.api/internal/models"
)

type IUserRepository interface {
	CreateUser(user *models.User) error
	CreateUserTx(tx *gorm.DB, user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	BeginTransaction() *gorm.DB
	UpdateUser(user *models.User) error
}
