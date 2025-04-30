package main

import (
	"github.com/sorawaslocked/ap2final_movie_service/config"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Print(cfg)
}
