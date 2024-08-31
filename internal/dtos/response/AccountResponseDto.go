package response

import "github.com/shopspring/decimal"

type AccountResponseDto struct {
	AccountNumber string          `json:"accountNumber"`
	AccountType   string          `json:"accountType"`
	Balance       decimal.Decimal `json:"balance"`
	Currency      string          `json:"currency"`
	Status        string          `json:"status"`
}
