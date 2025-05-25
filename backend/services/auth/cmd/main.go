package main

import (
	"fmt"
	"log"

	"github.com/keshvan/car-rental-platform/backend/pkg/config"
	"github.com/keshvan/car-rental-platform/backend/services/auth/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %w", err)
	}

	fmt.Println(cfg.AccessTTL, cfg.RefreshTTL)

	app.Run(cfg)
}
