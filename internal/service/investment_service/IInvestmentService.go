package investment_service

import (
	"github.com/google/uuid"
	"grip.app.api/internal/dtos/request"
	"grip.app.api/internal/dtos/response"
)

type IInvestmentService interface {
	CreateInvestment(req *request.CreateInvestmentRequest) (*response.InvestmentResponseDto, error)
	GetInvestmentByID(id uuid.UUID) (*response.InvestmentResponseDto, error)
	GetUserInvestmentPortfolio(userID uuid.UUID) (*response.UserInvestmentPortfolioResponse, error)
	UpdateInvestment(id uuid.UUID, req *request.UpdateInvestmentRequest) (*response.InvestmentResponseDto, error)
	DeleteInvestment(id uuid.UUID) error
}
