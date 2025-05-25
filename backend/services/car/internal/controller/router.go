package controller

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/keshvan/car-rental-platform/backend/pkg/jwt"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/controller/middleware"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/usecase"
)

func SetRoutes(engine *gin.Engine, carUsecase usecase.CarUsecase, reviewUsecase usecase.ReviewUsecase, rentUsecase usecase.RentUsecase, balanceUsecase usecase.BalanceUsecase, jwt *jwt.JWT) {
	carHandler := &CarHandler{carUsecase: carUsecase, reviewUsecase: reviewUsecase}
	rentHandler := &RentHandler{usecase: rentUsecase}
	balanceHandler := &BalanceHandler{usecase: balanceUsecase}
	reviewHandler := &ReviewHandler{reviewUsecase: reviewUsecase}
	auth := middleware.NewAuthMiddleware(jwt)

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:4200"},
		AllowMethods:     []string{"POST", "GET", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	engine.GET("/brands", carHandler.GetBrands)

	cars := engine.Group("/cars")
	{
		cars.GET("", carHandler.GetAll)
		cars.GET("/:id", carHandler.GetCarWithReviews)

		carAdmin := cars.Group("", auth.Auth(), middleware.RequireAdmin())
		{
			carAdmin.DELETE("/:id", carHandler.Delete)
			carAdmin.POST("", carHandler.NewCar)
			carAdmin.PATCH("/:id", carHandler.UpdateCar)
		}

	}

	rents := engine.Group("/rents", auth.Auth())
	{
		rents.POST("", rentHandler.NewRent)
		rents.GET("/me", rentHandler.GetMyRents)
		rents.PATCH("/complete/:id", rentHandler.CompleteRent)
		rents.GET("", middleware.RequireAdmin(), rentHandler.GetAllRents)
		rents.PATCH("/cancel/:id", middleware.RequireAdmin(), rentHandler.CancelRent)
		rents.DELETE("/:id", middleware.RequireAdmin(), rentHandler.DeleteRent)
	}

	balance := engine.Group("/balance", auth.Auth())
	{
		balance.PATCH("/", balanceHandler.UpdateBalance)
	}

	reviews := engine.Group("/reviews", auth.Auth(), middleware.RequireAdmin())
	{
		reviews.GET("", reviewHandler.GetUnverified)
		reviews.PATCH("/:id", reviewHandler.VerifyReview)
		reviews.DELETE("/:id", reviewHandler.Delete)
	}
}
