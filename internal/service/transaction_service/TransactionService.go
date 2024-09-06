package transaction_service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"grip.app.api/internal/dtos/request"
	"grip.app.api/internal/dtos/response"
	email_send_config2 "grip.app.api/internal/middleware/email_send_config"
	"grip.app.api/internal/models"
	"grip.app.api/internal/repository/account_repo"
	"grip.app.api/internal/repository/transaction_repo"
	"grip.app.api/internal/repository/user_repo"
	"time"
)

type TransactionService struct {
	transactionRepo transaction_repo.ITransactionRepository
	accountRepo     account_repo.IAccountRepository
	userRepo        user_repo.IUserRepository
}

func NewTransactionService(transactionRepo transaction_repo.ITransactionRepository, accountRepo account_repo.IAccountRepository, userRepo user_repo.IUserRepository) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
		accountRepo:     accountRepo,
		userRepo:        userRepo,
	}
}

func (s *TransactionService) Deposit(req request.DepositRequest) (*response.TransactionResponseDto, error) {
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, err
	}
	accountID, err := uuid.Parse(req.AccountID)
	if err != nil {
		return nil, err
	}
	amount := decimal.NewFromFloat(req.Amount)

	account, err := s.accountRepo.GetByID(accountID)
	if err != nil {
		return nil, err
	}

	newBalance := account.Balance.Add(amount)
	err = s.accountRepo.UpdateBalance(accountID, newBalance)
	if err != nil {
		return nil, err
	}

	transaction := &models.Transaction{
		AccountID:   accountID,
		UserID:      userID,
		Type:        models.Deposit,
		Amount:      amount,
		Currency:    req.Currency,
		Description: req.Description,
		Category:    req.Category,
		Status:      "COMPLETED",
		Timestamp:   time.Now(),
	}

	err = s.transactionRepo.Create(transaction)
	if err != nil {
		return nil, err
	}

	// Send email notification
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	err = email_send_config2.SendTransactionEmail(
		user.Email,
		user.FirstName,
		user.LastName,
		account.AccountNumber,
		email_send_config2.Deposit,
		amount,
		req.Currency,
		req.Description,
	)
	if err != nil {
		// Log the error, but don't return it as the deposit was successful
		fmt.Printf("Failed to send deposit confirmation email: %v\n", err)
	}

	return toTransactionResponseDto(transaction), nil
}

func (s *TransactionService) Withdraw(req request.WithdrawRequest) (*response.TransactionResponseDto, error) {
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, err
	}
	accountID, err := uuid.Parse(req.AccountID)
	if err != nil {
		return nil, err
	}
	amount := decimal.NewFromFloat(req.Amount)

	account, err := s.accountRepo.GetByID(accountID)
	if err != nil {
		return nil, err
	}

	if account.Balance.LessThan(amount) {
		return nil, errors.New("insufficient funds")
	}

	newBalance := account.Balance.Sub(amount)
	err = s.accountRepo.UpdateBalance(accountID, newBalance)
	if err != nil {
		return nil, err
	}

	transaction := &models.Transaction{
		AccountID:   accountID,
		UserID:      userID,
		Type:        models.Withdrawal,
		Amount:      amount.Neg(),
		Currency:    req.Currency,
		Description: req.Description,
		Category:    req.Category,
		Status:      "COMPLETED",
		Timestamp:   time.Now(),
	}

	err = s.transactionRepo.Create(transaction)
	if err != nil {
		return nil, err
	}

	// Send email notification
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	err = email_send_config2.SendTransactionEmail(
		user.Email,
		user.FirstName,
		user.LastName,
		account.AccountNumber,
		email_send_config2.Withdrawal,
		amount.Neg(),
		req.Currency,
		req.Description,
	)
	if err != nil {
		// Log the error, but don't return it as the withdrawal was successful
		fmt.Printf("Failed to send withdrawal confirmation email: %v\n", err)
	}

	return toTransactionResponseDto(transaction), nil
}

func (s *TransactionService) Transfer(req request.TransferRequest) (*request.TransferResponse, error) {
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, err
	}
	fromAccountID, err := uuid.Parse(req.FromAccountID)
	if err != nil {
		return nil, err
	}
	toAccountID, err := uuid.Parse(req.ToAccountID)
	if err != nil {
		return nil, err
	}
	amount := decimal.NewFromFloat(req.Amount)

	fromAccount, err := s.accountRepo.GetByID(fromAccountID)
	if err != nil {
		return nil, err
	}

	toAccount, err := s.accountRepo.GetByID(toAccountID)
	if err != nil {
		return nil, err
	}

	if fromAccount.Balance.LessThan(amount) {
		return nil, errors.New("insufficient funds")
	}

	newFromBalance := fromAccount.Balance.Sub(amount)
	newToBalance := toAccount.Balance.Add(amount)

	err = s.accountRepo.UpdateBalance(fromAccountID, newFromBalance)
	if err != nil {
		return nil, err
	}

	err = s.accountRepo.UpdateBalance(toAccountID, newToBalance)
	if err != nil {
		// Rollback the first update if the second fails
		_ = s.accountRepo.UpdateBalance(fromAccountID, fromAccount.Balance)
		return nil, err
	}

	fromTransaction := &models.Transaction{
		AccountID:   fromAccountID,
		UserID:      userID,
		Type:        models.Transfer,
		Amount:      amount.Neg(),
		Currency:    req.Currency,
		Description: req.Description,
		Category:    req.Category,
		Status:      "COMPLETED",
		Timestamp:   time.Now(),
	}

	toTransaction := &models.Transaction{
		AccountID:   toAccountID,
		UserID:      userID,
		Type:        models.Transfer,
		Amount:      amount,
		Currency:    req.Currency,
		Description: req.Description,
		Category:    req.Category,
		Status:      "COMPLETED",
		Timestamp:   time.Now(),
	}

	err = s.transactionRepo.Create(fromTransaction)
	if err != nil {
		return nil, err
	}

	err = s.transactionRepo.Create(toTransaction)
	if err != nil {
		return nil, err
	}

	// Send email notifications
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	err = email_send_config2.SendTransactionEmail(
		user.Email,
		user.FirstName,
		user.LastName,
		fromAccount.AccountNumber,
		email_send_config2.Transfer,
		amount.Neg(),
		req.Currency,
		req.Description,
	)
	if err != nil {
		fmt.Printf("Failed to send transfer confirmation email to sender: %v\n", err)
	}

	return &request.TransferResponse{
		FromTransaction: *fromTransaction,
		ToTransaction:   *toTransaction,
	}, nil
}

func (s *TransactionService) ChargeFee(userID, accountID uuid.UUID, amount decimal.Decimal, currency, description, category string) (*response.TransactionResponseDto, error) {
	account, err := s.accountRepo.GetByID(accountID)
	if err != nil {
		return nil, err
	}

	if account.Balance.LessThan(amount) {
		return nil, errors.New("insufficient funds")
	}

	newBalance := account.Balance.Sub(amount)
	err = s.accountRepo.UpdateBalance(accountID, newBalance)
	if err != nil {
		return nil, err
	}

	transaction := &models.Transaction{
		AccountID:   accountID,
		UserID:      userID,
		Type:        models.FeeCharged,
		Amount:      amount.Neg(),
		Currency:    currency,
		Description: description,
		Category:    category,
		Status:      "COMPLETED",
		Timestamp:   time.Now(),
	}

	err = s.transactionRepo.Create(transaction)
	if err != nil {
		return nil, err
	}

	// Send email notification
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	err = email_send_config2.SendTransactionEmail(
		user.Email,
		user.FirstName,
		user.LastName,
		account.AccountNumber,
		email_send_config2.FeeCharged,
		amount.Neg(),
		currency,
		description,
	)
	if err != nil {
		// Log the error, but don't return it as the fee charge was successful
		fmt.Printf("Failed to send fee charge notification email: %v\n", err)
	}

	return toTransactionResponseDto(transaction), nil
}

func (s *TransactionService) SendPayment(req request.TransferRequest) (*request.TransferResponse, error) {
	// For now, we'll implement SendPayment as identical to Transfer
	// In a real-world scenario, you might want to add additional logic specific to payments
	transferResponse, err := s.Transfer(req)
	if err != nil {
		return nil, err
	}

	// Send additional email notification for payment
	userID, _ := uuid.Parse(req.UserID)
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	fromAccountID, _ := uuid.Parse(req.FromAccountID)
	fromAccount, err := s.accountRepo.GetByID(fromAccountID)
	if err != nil {
		return nil, err
	}

	err = email_send_config2.SendTransactionEmail(
		user.Email,
		user.FirstName,
		user.LastName,
		fromAccount.AccountNumber,
		email_send_config2.Transfer, // You might want to create a new type for Payment in the email config
		decimal.NewFromFloat(req.Amount).Neg(),
		req.Currency,
		"Payment: "+req.Description,
	)
	if err != nil {
		fmt.Printf("Failed to send payment confirmation email: %v\n", err)
	}

	return transferResponse, nil
}

func toTransactionResponseDto(t *models.Transaction) *response.TransactionResponseDto {
	return &response.TransactionResponseDto{
		ID:          t.ID.String(),
		AccountID:   t.AccountID.String(),
		UserID:      t.UserID.String(),
		Type:        string(t.Type),
		Amount:      t.Amount.InexactFloat64(),
		Currency:    t.Currency,
		Description: t.Description,
		Category:    t.Category,
		Status:      t.Status,
		Timestamp:   t.Timestamp,
	}
}
