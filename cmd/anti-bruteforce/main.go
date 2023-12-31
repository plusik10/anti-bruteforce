package main

import (
	"log"

	"github.com/plusik10/anti-bruteforce/config"
	"github.com/plusik10/anti-bruteforce/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("%s", err)
	}
	app.Run(cfg)
}
