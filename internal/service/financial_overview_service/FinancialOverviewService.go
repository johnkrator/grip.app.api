package financial_overview_service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"grip.app.api/internal/dtos/response"
	"grip.app.api/internal/repository/financial_overview_repo"
	"grip.app.api/utils"
)

type FinancialOverviewService struct {
	repo financial_overview_repo.IFinancialOverviewRepository
}

func NewFinancialOverviewService(repo financial_overview_repo.IFinancialOverviewRepository) *FinancialOverviewService {
	return &FinancialOverviewService{repo: repo}
}

func (s *FinancialOverviewService) GetUserFinancialOverview(userID uuid.UUID) (*response.FinancialOverviewResponseDto, error) {
	user, account, loans, investments, err := s.repo.GetUserFinancialOverview(userID)
	if err != nil {
		if errors.Is(err, utils.ErrUserNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("error fetching financial overview: " + err.Error())
	}

	if user == nil {
		return nil, errors.New("user data is nil")
	}

	totalLoanAmount := decimal.Zero
	for _, loan := range loans {
		totalLoanAmount = totalLoanAmount.Add(loan.Amount)
	}

	totalInvestmentValue := decimal.Zero
	for _, investment := range investments {
		totalInvestmentValue = totalInvestmentValue.Add(investment.CurrentValue)
	}

	return &response.FinancialOverviewResponseDto{
		User: response.UserResponseDto{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		},
		Account: response.AccountResponseDto{
			AccountNumber: account.AccountNumber,
			Balance:       account.Balance,
			Currency:      string(account.Currency),
		},
		TotalLoanAmount:      totalLoanAmount,
		TotalInvestmentValue: totalInvestmentValue,
		NetWorth:             account.Balance.Add(totalInvestmentValue).Sub(totalLoanAmount),
	}, nil
}
