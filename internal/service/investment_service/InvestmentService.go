package investment_service

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"grip.app.api/internal/dtos/request"
	"grip.app.api/internal/dtos/response"
	"grip.app.api/internal/models"
	"grip.app.api/internal/repository/investment_repo"
)

type InvestmentService struct {
	repo investment_repo.IInvestmentRepository
}

func NewInvestmentService(repo investment_repo.IInvestmentRepository) *InvestmentService {
	return &InvestmentService{repo: repo}
}

func (s *InvestmentService) CreateInvestment(req *request.CreateInvestmentRequest) (*response.InvestmentResponseDto, error) {
	investment := &models.Investment{
		UserID:       req.UserID,
		Type:         req.Type,
		Amount:       req.Amount,
		PurchaseDate: req.PurchaseDate,
		CurrentValue: req.Amount, // Initially set to the purchase amount
		Status:       "Active",   // Default status
	}
	err := s.repo.Create(investment)
	if err != nil {
		return nil, err
	}
	return toInvestmentResponse(investment), nil
}

func (s *InvestmentService) GetInvestmentByID(id uuid.UUID) (*response.InvestmentResponseDto, error) {
	investment, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return toInvestmentResponse(investment), nil
}

func (s *InvestmentService) GetUserInvestmentPortfolio(userID uuid.UUID) (*response.UserInvestmentPortfolioResponse, error) {
	investments, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	portfolioResponse := &response.UserInvestmentPortfolioResponse{
		UserID:      userID,
		Investments: make([]response.InvestmentResponseDto, len(investments)),
	}

	for i, investment := range investments {
		portfolioResponse.Investments[i] = *toInvestmentResponse(investment)
		portfolioResponse.TotalPortfolioValue = portfolioResponse.TotalPortfolioValue.Add(investment.CurrentValue)
	}

	return portfolioResponse, nil
}

func (s *InvestmentService) UpdateInvestment(id uuid.UUID, req *request.UpdateInvestmentRequest) (*response.InvestmentResponseDto, error) {
	investment, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields from the request
	investment.Type = req.Type
	investment.Amount = req.Amount
	investment.CurrentValue = req.Amount
	investment.Status = "Active"

	err = s.repo.Update(investment)
	if err != nil {
		return nil, err
	}

	return toInvestmentResponse(investment), nil
}

func (s *InvestmentService) DeleteInvestment(id uuid.UUID) error {
	// First, check if the investment exists
	_, err := s.repo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("investment not found")
		}
		return err
	}

	// If the investment exists, proceed with deletion
	err = s.repo.Delete(id)
	if err != nil {
		return err
	}

	// Verify deletion
	_, err = s.repo.GetByID(id)
	if err == nil {
		return fmt.Errorf("investment still exists after deletion")
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	return nil
}

func toInvestmentResponse(investment *models.Investment) *response.InvestmentResponseDto {
	return &response.InvestmentResponseDto{
		ID:           investment.ID,
		UserID:       investment.UserID,
		Type:         investment.Type,
		Amount:       investment.Amount,
		PurchaseDate: investment.PurchaseDate,
		CurrentValue: investment.CurrentValue,
		Status:       investment.Status,
	}
}
