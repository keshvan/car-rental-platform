package controller

import (
	"errors"
	"fmt"
	"net/http"
	"net/mail"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/controller/request"
	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/repo"
	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/usecase"
)

var (
	ErrInvalidEmail = errors.New("invalid email")
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

	if !validateEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidEmail.Error()})
	}

	res, err := h.usecase.Register(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, repo.ErrDuplicateEmail) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		}
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

	if !validateEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidEmail.Error()})
	}

	cookieName := "refresh_token"
	if strings.Contains(c.Request.Header.Get("Origin"), ":4200") {
		cookieName = "refresh_token_admin"
	}

	oldRefreshToken, err := c.Cookie(cookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			oldRefreshToken = ""
		} else {
			oldRefreshToken = ""
		}
	}

	res, err := h.usecase.Login(c.Request.Context(), req.Email, req.Password, oldRefreshToken, cookieName)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidEmailOrPassword) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, usecase.ErrInvalidRole) {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie(cookieName, res.Tokens.RefreshToken, 3600*24*30, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"user": res.User, "access_token": res.Tokens.AccessToken})
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	cookieName := "refresh_token"
	if strings.Contains(c.Request.Header.Get("Origin"), ":4200") {
		cookieName = "refresh_token_admin"
	}

	fmt.Println("cookieName", cookieName)

	refreshToken, err := c.Cookie(cookieName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token required"})
		return
	}

	tokens, err := h.usecase.Refresh(c.Request.Context(), refreshToken)
	if err != nil {
		fmt.Println("err", err)
		c.SetCookie(cookieName, "", -1, "/", "", false, true)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired refresh token"})
		return
	}

	c.SetCookie(cookieName, tokens.RefreshToken, 3600*24*30, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"access_token": tokens.AccessToken})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err == nil && refreshToken != "" {
		_ = h.usecase.Logout(c.Request.Context(), refreshToken)
	}

	cookieName := "refresh_token"
	if strings.Contains(c.Request.Host, ":4200") {
		cookieName = "refresh_token_admin"
	}

	c.SetCookie(cookieName, "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}

func (h *AuthHandler) CheckSession(c *gin.Context) {
	cookieName := "refresh_token"
	if strings.Contains(c.Request.Header.Get("Origin"), ":4200") {
		cookieName = "refresh_token_admin"
	}

	refreshToken, err := c.Cookie(cookieName)
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

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
