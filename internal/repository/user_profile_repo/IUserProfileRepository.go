package user_profile_repo

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"grip.app.api/internal/models"
)

type IUserProfileRepository interface {
	CreateUserProfileTx(tx *gorm.DB, profile *models.UserProfile) error
	GetUserProfileByUserID(userID uuid.UUID) (*models.UserProfile, error)
	UpdateUserProfile(profile *models.UserProfile) error
}
