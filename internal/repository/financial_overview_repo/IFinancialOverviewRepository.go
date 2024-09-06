package financial_overview_repo

import (
	"github.com/google/uuid"
	"grip.app.api/internal/models"
)

type IFinancialOverviewRepository interface {
	GetUserFinancialOverview(userID uuid.UUID) (*models.User, *models.Account, []*models.Loan, []*models.Investment, error)
}
