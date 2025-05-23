package main

import (
	"log"

	"github.com/keshvan/car-rental-platform/backend/pkg/config"
	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %w", err)
	}

	app.Run(cfg)
}
