package main

import (
	"core/internal/config"
	"core/internal/database"
	"core/internal/router"
	"log"
)

func main() {
	cfg := config.Load()

	_, err := database.Init(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := router.Setup()

	if err := r.Run(":" + cfg.App.Port); err != nil {
		log.Fatal(err)
	}
}
