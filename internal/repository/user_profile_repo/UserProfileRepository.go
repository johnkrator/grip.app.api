package user_profile_repo

import (
	"github.com/google/uuid"
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

func (r *UserProfileRepository) GetUserProfileByUserID(userID uuid.UUID) (*models.UserProfile, error) {
	var profile models.UserProfile
	err := r.db.Where("user_id = ?", userID).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *UserProfileRepository) UpdateUserProfile(profile *models.UserProfile) error {
	return r.db.Save(profile).Error
}
