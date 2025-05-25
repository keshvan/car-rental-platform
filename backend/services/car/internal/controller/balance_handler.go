package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/controller/middleware"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/entity"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/usecase"
)

type BalanceHandler struct {
	usecase usecase.BalanceUsecase
}

func (h *BalanceHandler) UpdateBalance(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var balance entity.Balance
	if err := c.ShouldBindJSON(&balance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	balance.UserID = userID

	if err := h.usecase.UpdateBalance(c.Request.Context(), &balance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, balance)
}
