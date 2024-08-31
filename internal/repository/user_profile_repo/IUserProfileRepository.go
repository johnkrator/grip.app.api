package user_profile_repo

import (
	"gorm.io/gorm"
	"grip.app.api/internal/models"
)

type IUserProfileRepository interface {
	CreateUserProfileTx(tx *gorm.DB, profile *models.UserProfile) error
}
