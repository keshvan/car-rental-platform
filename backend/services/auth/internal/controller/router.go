package controller

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/keshvan/car-rental-platform/backend/pkg/jwt"
	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/controller/middleware"
	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/usecase"
)

func SetRoutes(engine *gin.Engine, authUsecase usecase.AuthUsecase, userUsecase usecase.UserUsecase, jwt *jwt.JWT) {
	h := &AuthHandler{authUsecase}
	userH := &UserHandler{userUsecase}
	auth := middleware.NewAuthMiddleware(jwt)

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:4200"},
		AllowMethods:     []string{"POST", "GET", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	engine.POST("/register", h.Register)
	engine.POST("/login", h.Login)
	engine.POST("/refresh", h.Refresh)
	engine.POST("/logout", h.Logout)
	engine.GET("/check-session", h.CheckSession)

	userRoutes := engine.Group("/users", auth.Auth(), middleware.RequireAdmin())
	{
		userRoutes.GET("", userH.GetAllUsers)
		userRoutes.PATCH("/:id", userH.UpdateUserRole)
		userRoutes.DELETE("/:id", userH.DeleteUser)
	}
}
