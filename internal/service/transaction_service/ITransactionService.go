package transaction_service

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"grip.app.api/internal/dtos/request"
	"grip.app.api/internal/dtos/response"
)

type ITransactionService interface {
	Deposit(req request.DepositRequest) (*response.TransactionResponseDto, error)
	Withdraw(req request.WithdrawRequest) (*response.TransactionResponseDto, error)
	Transfer(req request.TransferRequest) (*request.TransferResponse, error)
	SendPayment(req request.TransferRequest) (*request.TransferResponse, error)
	ChargeFee(userID, accountID uuid.UUID, amount decimal.Decimal, currency, description, category string) (*response.TransactionResponseDto, error)
}
