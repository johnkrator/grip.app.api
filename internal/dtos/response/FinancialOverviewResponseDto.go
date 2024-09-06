package response

import "github.com/shopspring/decimal"

type FinancialOverviewResponseDto struct {
	User                 UserResponseDto         `json:"user"`
	Account              AccountResponseDto      `json:"account"`
	TotalLoanAmount      decimal.Decimal         `json:"totalLoanAmount"`
	TotalInvestmentValue decimal.Decimal         `json:"totalInvestmentValue"`
	NetWorth             decimal.Decimal         `json:"netWorth"`
	Loans                []LoanResponseDto       `json:"loans"`
	Investments          []InvestmentResponseDto `json:"investments"`
}
