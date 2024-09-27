package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"grip.app.api/internal/dtos/request"
	"grip.app.api/internal/service/investment_service"
	"net/http"
)

type InvestmentController struct {
	service investment_service.IInvestmentService
}

func NewInvestmentController(service investment_service.IInvestmentService) *InvestmentController {
	return &InvestmentController{service: service}
}

func (c *InvestmentController) CreateInvestment(ctx *gin.Context) {
	var req request.CreateInvestmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.service.CreateInvestment(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *InvestmentController) GetInvestmentByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid investment ID"})
		return
	}

	investment, err := c.service.GetInvestmentByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Investment not found"})
		return
	}

	ctx.JSON(http.StatusOK, investment)
}

func (c *InvestmentController) GetUserInvestmentPortfolio(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	portfolio, err := c.service.GetUserInvestmentPortfolio(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, portfolio)
}

func (c *InvestmentController) UpdateInvestment(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid investment ID"})
		return
	}

	var req request.UpdateInvestmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.service.UpdateInvestment(id, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *InvestmentController) DeleteInvestment(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid investment ID"})
		return
	}

	err = c.service.DeleteInvestment(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Investment deleted successfully"})
}
