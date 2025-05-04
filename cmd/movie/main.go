package main

import (
	"github.com/sorawaslocked/ap2final_movie_service/internal/config"
	"log"
)

func main() {
	// TODO: load config
	cfg := config.MustLoad()
	log.Println(cfg)

	// TODO: setup logger

	// TODO: initialize app and run it
}
