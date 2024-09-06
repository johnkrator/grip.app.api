package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"grip.app.api/internal/service/financial_overview_service"
	"net/http"
)

type FinancialOverviewController struct {
	service financial_overview_service.IFinancialOverviewService
}

func NewFinancialOverviewController(service financial_overview_service.IFinancialOverviewService) *FinancialOverviewController {
	return &FinancialOverviewController{service: service}
}

// GetUserFinancialOverview godoc
// @Summary Get user financial overview
// @Description Get the authenticated user's financial overview
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} response.FinancialOverviewResponseDto
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id}/financial-overview [get]
func (c *FinancialOverviewController) GetUserFinancialOverview(ctx *gin.Context) {
	userIDStr := ctx.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	overview, err := c.service.GetUserFinancialOverview(userID)
	if err != nil {
		if err.Error() == "user not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch financial overview"})
		}
		return
	}

	ctx.JSON(http.StatusOK, overview)
}
