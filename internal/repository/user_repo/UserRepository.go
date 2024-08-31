package user_repo

import (
	"gorm.io/gorm"
	"grip.app.api/internal/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
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
