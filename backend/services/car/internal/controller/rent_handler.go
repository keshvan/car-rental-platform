package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/controller/middleware"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/entity"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/repo"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/usecase"
)

type RentHandler struct {
	usecase usecase.RentUsecase
}

func (h *RentHandler) NewRent(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var rent entity.Rent
	if err := c.ShouldBindJSON(&rent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rent.UserID = userID

	if err := h.usecase.NewRent(c.Request.Context(), &rent); err != nil {
		switch err {
		case usecase.ErrNotEnoughBalance:
			c.JSON(http.StatusPaymentRequired, gin.H{"error": err.Error()})
		case usecase.ErrCarNotAvailable:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "rent created successfully"})
}

func (h *RentHandler) GetAllRents(c *gin.Context) {
	rents, err := h.usecase.GetAllRents(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"rents": rents})
}

func (h *RentHandler) GetMyRents(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	rents, err := h.usecase.GetRentsByUserID(c.Request.Context(), userID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"rents": rents})
}

func (h *RentHandler) CompleteRent(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	role, _ := middleware.GetRoleFromContext(c)
	idParam := c.Param("id")
	rentId, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid rent id"})
	}

	var review entity.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.CompleteRent(c.Request.Context(), rentId, &review, userID, role); err != nil {
		fmt.Println(err)
		switch err {
		case usecase.ErrNoAccess:
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		case repo.ErrRentNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "rent successfully completed"})
}

func (h *RentHandler) CancelRent(c *gin.Context) {
	idParam := c.Param("id")
	rentId, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid rent id"})
	}

	if err := h.usecase.Cancel(c.Request.Context(), rentId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "rent successfully cancelled"})
}

func (h *RentHandler) DeleteRent(c *gin.Context) {
	idParam := c.Param("id")
	rentId, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid rent id"})
	}

	if err := h.usecase.Delete(c.Request.Context(), rentId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "rent successfully deleted"})
}
