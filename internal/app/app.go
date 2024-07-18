package app

import (
	"github.com/sudo-odner/LifeManagerBack/config"
	"github.com/sudo-odner/LifeManagerBack/internal/controller/http/httpHandler"
	"github.com/sudo-odner/LifeManagerBack/internal/controller/http/httpServer"
	"github.com/sudo-odner/LifeManagerBack/internal/logger"
	"github.com/sudo-odner/LifeManagerBack/internal/usecase"
)

func Run(cfg *config.Config) {
	// TODO: Write logging
	log := logger.New(cfg.Logger)
	// TODO: Write repository

	// TODO: Write usecase
	u := usecase.New()
	// Init HTTP Server - controller
	// Init handler
	handler := httpHandler.Create(log, u)
	// Init server
	httpServer.New(handler)
}
