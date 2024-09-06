package request

import "grip.app.api/internal/models"

type DepositRequest struct {
	UserID      string  `json:"userId" binding:"required"`
	AccountID   string  `json:"accountId" binding:"required"`
	Amount      float64 `json:"amount" binding:"required,gt=0"`
	Currency    string  `json:"currency" binding:"required"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}

type WithdrawRequest struct {
	UserID      string  `json:"userId" binding:"required"`
	AccountID   string  `json:"accountId" binding:"required"`
	Amount      float64 `json:"amount" binding:"required,gt=0"`
	Currency    string  `json:"currency" binding:"required"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}

type TransferRequest struct {
	UserID        string  `json:"userId" binding:"required"`
	FromAccountID string  `json:"fromAccountId" binding:"required"`
	ToAccountID   string  `json:"toAccountId" binding:"required"`
	Amount        float64 `json:"amount" binding:"required,gt=0"`
	Currency      string  `json:"currency" binding:"required"`
	Description   string  `json:"description"`
	Category      string  `json:"category"`
}

type TransferResponse struct {
	FromTransaction models.Transaction `json:"fromTransaction"`
	ToTransaction   models.Transaction `json:"toTransaction"`
}
