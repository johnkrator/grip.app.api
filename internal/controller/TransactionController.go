package controller

import (
	"github.com/gin-gonic/gin"
	request2 "grip.app.api/internal/dtos/request"
	"grip.app.api/internal/service/transaction_service"
	"net/http"
)

type TransactionController struct {
	service transaction_service.ITransactionService
}

func NewTransactionController(service transaction_service.ITransactionService) *TransactionController {
	return &TransactionController{service: service}
}

// Deposit godoc
// @Summary Deposit funds
// @Description Deposit funds into a user's account
// @Tags transactions
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body request.DepositRequest true "Deposit request"
// @Success 200 {object} response.TransactionResponseDto
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /transactions/deposit [post]
func (c *TransactionController) Deposit(ctx *gin.Context) {
	var req request2.DepositRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := c.service.Deposit(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tx)
}

// Withdraw godoc
// @Summary Withdraw funds
// @Description Withdraw funds from a user's account
// @Tags transactions
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body request.WithdrawRequest true "Withdraw request"
// @Success 200 {object} response.TransactionResponseDto
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /transactions/withdraw [post]
func (c *TransactionController) Withdraw(ctx *gin.Context) {
	var req request2.WithdrawRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := c.service.Withdraw(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tx)
}

// Transfer godoc
// @Summary Transfer funds
// @Description Transfer funds between two accounts
// @Tags transactions
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body request.TransferRequest true "Transfer request"
// @Success 200 {object} request.TransferResponse
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /transactions/transfer [post]
func (c *TransactionController) Transfer(ctx *gin.Context) {
	var req request2.TransferRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transferResponse, err := c.service.Transfer(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transferResponse)
}

//package controller
//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/google/uuid"
//	"github.com/shopspring/decimal"
//	"grip.app.api/internal/service/transaction_service"
//	"net/http"
//)
//
//type TransactionController struct {
//	service transaction_service.ITransactionService
//}
//
//func NewTransactionController(service transaction_service.ITransactionService) *TransactionController {
//	return &TransactionController{service: service}
//}
//
//// Deposit godoc
//// @Summary Deposit funds
//// @Description Deposit funds into a user's account
//// @Tags transactions
//// @Accept json
//// @Produce json
//// @Param Authorization header string true "Bearer token"
//// @Param request body DepositRequest true "Deposit request"
//// @Success 200 {object} models.Transaction
//// @Failure 400 {object} gin.H
//// @Failure 401 {object} gin.H
//// @Failure 500 {object} gin.H
//// @Router /transactions/deposit [post]
//func (c *TransactionController) Deposit(ctx *gin.Context) {
//	var req struct {
//		UserID      string  `json:"userId" binding:"required"`
//		AccountID   string  `json:"accountId" binding:"required"`
//		Amount      float64 `json:"amount" binding:"required,gt=0"`
//		Currency    string  `json:"currency" binding:"required"`
//		Description string  `json:"description"`
//		Category    string  `json:"category"`
//	}
//
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	userID, err := uuid.Parse(req.UserID)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
//		return
//	}
//
//	accountID, err := uuid.Parse(req.AccountID)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
//		return
//	}
//
//	amount := decimal.NewFromFloat(req.Amount)
//	tx, err := c.service.Deposit(userID, accountID, amount, req.Currency, req.Description, req.Category)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, tx)
//}
//
//// Withdraw godoc
//// @Summary Withdraw funds
//// @Description Withdraw funds from a user's account
//// @Tags transactions
//// @Accept json
//// @Produce json
//// @Param Authorization header string true "Bearer token"
//// @Param request body WithdrawRequest true "Withdraw request"
//// @Success 200 {object} models.Transaction
//// @Failure 400 {object} gin.H
//// @Failure 401 {object} gin.H
//// @Failure 500 {object} gin.H
//// @Router /transactions/withdraw [post]
//func (c *TransactionController) Withdraw(ctx *gin.Context) {
//	var req struct {
//		UserID      string  `json:"userId" binding:"required"`
//		AccountID   string  `json:"accountId" binding:"required"`
//		Amount      float64 `json:"amount" binding:"required,gt=0"`
//		Currency    string  `json:"currency" binding:"required"`
//		Description string  `json:"description"`
//		Category    string  `json:"category"`
//	}
//
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	userID, err := uuid.Parse(req.UserID)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
//		return
//	}
//
//	accountID, err := uuid.Parse(req.AccountID)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
//		return
//	}
//
//	amount := decimal.NewFromFloat(req.Amount)
//	tx, err := c.service.Withdraw(userID, accountID, amount, req.Currency, req.Description, req.Category)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	ctx.JSON(http.StatusOK, tx)
//}
//
//// Transfer godoc
//// @Summary Transfer funds
//// @Description Transfer funds between two accounts
//// @Tags transactions
//// @Accept json
//// @Produce json
//// @Param Authorization header string true "Bearer token"
//// @Param request body TransferRequest true "Transfer request"
//// @Success 200 {object} TransferResponse
//// @Failure 400 {object} gin.H
//// @Failure 401 {object} gin.H
//// @Failure 500 {object} gin.H
//// @Router /transactions/transfer [post]
//func (c *TransactionController) Transfer(ctx *gin.Context) {
//	var req struct {
//		UserID        string  `json:"userId" binding:"required"`
//		FromAccountID string  `json:"fromAccountId" binding:"required"`
//		ToAccountID   string  `json:"toAccountId" binding:"required"`
//		Amount        float64 `json:"amount" binding:"required,gt=0"`
//		Currency      string  `json:"currency" binding:"required"`
//		Description   string  `json:"description"`
//		Category      string  `json:"category"`
//	}
//
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	userID, err := uuid.Parse(req.UserID)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
//		return
//	}
//
//	fromAccountID, err := uuid.Parse(req.FromAccountID)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid from account ID"})
//		return
//	}
//
//	toAccountID, err := uuid.Parse(req.ToAccountID)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid to account ID"})
//		return
//	}
//
//	amount := decimal.NewFromFloat(req.Amount)
//	fromTx, toTx, err := c.service.Transfer(userID, fromAccountID, toAccountID, amount, req.Currency, req.Description, req.Category)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{"fromTransaction": fromTx, "toTransaction": toTx})
//}
