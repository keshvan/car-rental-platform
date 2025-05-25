package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/controller/response"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/entity"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/usecase"
)

type CarHandler struct {
	carUsecase    usecase.CarUsecase
	reviewUsecase usecase.ReviewUsecase
}

func (h *CarHandler) GetAll(c *gin.Context) {
	cars, err := h.carUsecase.GetAll(c.Request.Context())
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cars": cars})
}

func (h *CarHandler) GetBrands(c *gin.Context) {
	brands, err := h.carUsecase.GetBrands(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"brands": brands})
}

func (h *CarHandler) GetCarWithReviews(c *gin.Context) {
	carIDParam := c.Param("id")
	carID, err := strconv.ParseInt(carIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid car ID"})
		return
	}

	car, err := h.carUsecase.GetByID(c.Request.Context(), carID)
	if err != nil {
		if errors.Is(err, usecase.ErrCarNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	reviews, err := h.reviewUsecase.GetByCarID(c.Request.Context(), car.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	carWithReviews := response.CarWithReviews{
		Car:     *car,
		Reviews: reviews,
	}

	c.JSON(http.StatusOK, carWithReviews)
}

func (h *CarHandler) Delete(c *gin.Context) {
	carIDParam := c.Param("id")
	carID, err := strconv.ParseInt(carIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid car ID"})
		return
	}

	if err := h.carUsecase.Delete(c.Request.Context(), carID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "car deleted successfully"})
}

func (h *CarHandler) NewCar(c *gin.Context) {
	var car entity.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.carUsecase.NewCar(c.Request.Context(), &car); err != nil {
		fmt.Println(err)
		if errors.Is(err, usecase.ErrBrandNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "car created successfully"})
}

func (h *CarHandler) UpdateCar(c *gin.Context) {
	carIDParam := c.Param("id")
	carID, err := strconv.ParseInt(carIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid car ID"})
		return
	}

	var car entity.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.carUsecase.Update(c.Request.Context(), carID, &car); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "car updated successfully"})
}
