package response

import "time"

type TransactionResponseDto struct {
	ID          string    `json:"id"`
	AccountID   string    `json:"accountId"`
	UserID      string    `json:"userId"`
	Type        string    `json:"type"`
	Amount      float64   `json:"amount"`
	Currency    string    `json:"currency"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Status      string    `json:"status"`
	Timestamp   time.Time `json:"timestamp"`
}
