package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/usecase"
)

type ReviewHandler struct {
	reviewUsecase usecase.ReviewUsecase
}

func (h *ReviewHandler) GetUnverified(c *gin.Context) {
	reviews, err := h.reviewUsecase.GetUnverified(c.Request.Context())
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}

func (h *ReviewHandler) VerifyReview(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}
	err = h.reviewUsecase.VerifyReview(c.Request.Context(), idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review verified successfully"})
}

func (h *ReviewHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}
	err = h.reviewUsecase.Delete(c.Request.Context(), idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}
