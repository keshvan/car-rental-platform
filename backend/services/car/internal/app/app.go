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
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/controller"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/repo"
	"github.com/keshvan/car-rental-platform/backend/services/car/internal/usecase"
)

func Run(cfg *config.Config) {
	fmt.Println(cfg)
	//Repositories
	db, err := database.New(cfg.DbPath)
	if err != nil {
		log.Fatalf("App - Run - database.New()")
	}
	defer db.Close()

	carRepo := repo.NewCarRepository(db)

	//JWT
	jwt := jwt.New(cfg.Secret, cfg.AccessTTL, cfg.RefreshTTL)

	//Usecase
	carUsecase := usecase.NewCarUsecase(carRepo)

	//Server
	httpServer := httpserver.New(cfg.Server)
	controller.SetRoutes(httpServer.Engine, *carUsecase, jwt)
	httpServer.Run()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt
}
