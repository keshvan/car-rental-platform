package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/usecase"
)

type CarHandler struct {
	usecase usecase.CarUsecase
}

func (h *CarHandler) GetAll(c *gin.Context) {
	cars, err := h.usecase.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cars": cars})
}
