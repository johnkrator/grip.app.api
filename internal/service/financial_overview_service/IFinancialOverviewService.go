package financial_overview_service

import (
	"github.com/google/uuid"
	"grip.app.api/internal/dtos/response"
)

type IFinancialOverviewService interface {
	GetUserFinancialOverview(userID uuid.UUID) (*response.FinancialOverviewResponseDto, error)
}
