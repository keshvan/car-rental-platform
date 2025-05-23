package controller

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/keshvan/car-rental-platform/backend/pkg/jwt"
	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/usecase"
)

func SetRoutes(engine *gin.Engine, usecase usecase.AuthUsecase, jwt *jwt.JWT) {
	h := &AuthHandler{usecase}

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	engine.POST("/register", h.Register)
	engine.POST("/login", h.Login)
	engine.POST("/refresh", h.Refresh)
	engine.POST("/logout", h.Logout)
	engine.GET("/check-session", h.CheckSession)
}
