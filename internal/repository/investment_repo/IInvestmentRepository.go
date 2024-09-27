package investment_repo

import (
	"github.com/google/uuid"
	"grip.app.api/internal/models"
)

type IInvestmentRepository interface {
	Create(investment *models.Investment) error
	GetByID(id uuid.UUID) (*models.Investment, error)
	GetByUserID(userID uuid.UUID) ([]*models.Investment, error)
	Update(investment *models.Investment) error
	Delete(id uuid.UUID) error
}
