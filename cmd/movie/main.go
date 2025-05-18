package main

import (
	"github.com/sorawaslocked/ap2final_base/pkg/logger"
	"github.com/sorawaslocked/ap2final_movie_service/internal/config"
)

func main() {
	// TODO: load config
	cfg := config.MustLoad()

	// TODO: setup logger
	log := logger.SetupLogger(cfg.Env)

	// TODO: initialize app and run it
	log.Info("initializing application")
}
