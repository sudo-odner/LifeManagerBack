package main

import (
	"github.com/sudo-odner/LifeManagerBack/config"
	"github.com/sudo-odner/LifeManagerBack/internal/app"
)

func main() {
	// Upload config
	cfg, err := config.New()
	if err != nil {
		panic("Config not load, server can't start")
	}
	// Run Application
	app.Run(cfg)
}
