package financial_overview_repo

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"grip.app.api/internal/models"
)

type FinancialOverviewRepository struct {
	db *gorm.DB
}

func NewFinancialOverviewRepository(db *gorm.DB) IFinancialOverviewRepository {
	return &FinancialOverviewRepository{db: db}
}

func (r *FinancialOverviewRepository) GetUserFinancialOverview(userID uuid.UUID) (*models.User, *models.Account, []*models.Loan, []*models.Investment, error) {
	var user models.User
	var account models.Account
	var loans []*models.Loan
	var investments []*models.Investment

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user, userID).Error; err != nil {
			return err
		}

		if err := tx.Where("user_id = ?", userID).First(&account).Error; err != nil {
			return err
		}

		if err := tx.Where("user_id = ?", userID).Find(&loans).Error; err != nil {
			return err
		}

		if err := tx.Where("user_id = ?", userID).Find(&investments).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, nil, nil, nil, err
	}

	return &user, &account, loans, investments, nil
}
