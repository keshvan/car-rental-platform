package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/controller/request"
	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/usecase"
)

type AuthHandler struct {
	usecase usecase.AuthUsecase
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.usecase.Register(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("refresh_token", res.Tokens.RefreshToken, 3600*24*30, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"user": res.User, "access_token": res.Tokens.AccessToken})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	oldRefreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		if err == http.ErrNoCookie {
			oldRefreshToken = ""
		} else {
			oldRefreshToken = ""
		}
	}

	res, err := h.usecase.Login(c.Request.Context(), req.Username, req.Password, oldRefreshToken)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidEmailOrPassword) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("refresh_token", res.Tokens.RefreshToken, 3600*24*30, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"user": res.User, "access_token": res.Tokens.AccessToken})
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token required"})
		return
	}

	tokens, err := h.usecase.Refresh(c.Request.Context(), refreshToken)
	if err != nil {
		c.SetCookie("refresh_token", "", -1, "/", "", false, true)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired refresh token"})
		return
	}

	c.SetCookie("refresh_token", tokens.RefreshToken, 3600*24*30, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"access_token": tokens.AccessToken})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err == nil && refreshToken != "" {
		_ = h.usecase.Logout(c.Request.Context(), refreshToken)
	}

	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}

func (h *AuthHandler) CheckSession(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"user": nil, "is_active": false})
		return
	}

	res, err := h.usecase.CheckSession(c.Request.Context(), refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"user": nil, "is_active": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": res.User, "is_active": res.IsActive})
}
