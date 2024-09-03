package user_repo

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"grip.app.api/internal/models"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) CreateUserTx(tx *gorm.DB, user *models.User) error {
	return tx.Create(user).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUserToken(userID uuid.UUID, token string, expiration time.Time) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"token":            token,
		"token_expiration": expiration,
	}).Error
}

func (r *UserRepository) GetUserByResetToken(token string) (*models.User, error) {
	var user models.User
	result := r.db.Where("reset_password_token = ?", token).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) SoftDeleteUser(id uuid.UUID) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Update("deleted_at", gorm.Expr("NOW()")).Error
}

func (r *UserRepository) GetAllUsers(page, pageSize int) ([]*models.User, int64, error) {
	var users []*models.User
	var totalCount int64

	offset := (page - 1) * pageSize

	err := r.db.Model(&models.User{}).Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Offset(offset).Limit(pageSize).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}
