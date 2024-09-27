package investment_repo

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"grip.app.api/internal/models"
)

type InvestmentRepository struct {
	db *gorm.DB
}

func NewInvestmentRepository(db *gorm.DB) *InvestmentRepository {
	return &InvestmentRepository{db: db}
}

func (r *InvestmentRepository) Create(investment *models.Investment) error {
	return r.db.Create(investment).Error
}

func (r *InvestmentRepository) GetByID(id uuid.UUID) (*models.Investment, error) {
	var investment models.Investment
	err := r.db.First(&investment, id).Error
	if err != nil {
		return nil, err
	}
	return &investment, nil
}

func (r *InvestmentRepository) GetByUserID(userID uuid.UUID) ([]*models.Investment, error) {
	var investments []*models.Investment
	err := r.db.Where("user_id = ?", userID).Find(&investments).Error
	return investments, err
}

func (r *InvestmentRepository) Update(investment *models.Investment) error {
	return r.db.Save(investment).Error
}

func (r *InvestmentRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Investment{}, id).Error
}
