package user_repo

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"grip.app.api/internal/models"
	"time"
)

type IUserRepository interface {
	CreateUser(user *models.User) error
	CreateUserTx(tx *gorm.DB, user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	BeginTransaction() *gorm.DB
	UpdateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	UpdateUserToken(userID uuid.UUID, token string, expiration time.Time) error
}
