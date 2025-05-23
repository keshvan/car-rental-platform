package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/keshvan/car-rental-platform/backend/pkg/config"
	"github.com/keshvan/car-rental-platform/backend/pkg/database"
	"github.com/keshvan/car-rental-platform/backend/pkg/httpserver"
	"github.com/keshvan/car-rental-platform/backend/pkg/jwt"
	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/controller"
	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/repo"
	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/usecase"
)

func Run(cfg *config.Config) {
	fmt.Println(cfg)
	//Repositories
	db, err := database.New(cfg.DbPath)
	if err != nil {
		log.Fatalf("App - Run - database.New()")
	}
	defer db.Close()

	userRepo := repo.NewUserRepo(db)
	tokenRepo := repo.NewTokenRepo(db)

	//JWT
	jwt := jwt.New(cfg.Secret, cfg.AccessTTL, cfg.RefreshTTL)

	//Usecase
	authUsecase := usecase.NewAuthUsecase(userRepo, tokenRepo, jwt)

	//Server
	httpServer := httpserver.New(cfg.Server)
	controller.SetRoutes(httpServer.Engine, authUsecase, jwt)
	httpServer.Run()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt
}
