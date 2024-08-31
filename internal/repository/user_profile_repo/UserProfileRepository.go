package user_profile_repo

import (
	"gorm.io/gorm"
	"grip.app.api/internal/models"
)

type UserProfileRepository struct {
	db *gorm.DB
}

func NewUserProfileRepository(db *gorm.DB) IUserProfileRepository {
	return &UserProfileRepository{db: db}
}

func (r *UserProfileRepository) CreateUserProfileTx(tx *gorm.DB, profile *models.UserProfile) error {
	return tx.Create(profile).Error
}
