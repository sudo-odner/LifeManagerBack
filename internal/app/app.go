package app

import (
	"github.com/sudo-odner/LifeManagerBack/internal/controller/http/httpHandler"
	"github.com/sudo-odner/LifeManagerBack/internal/controller/http/httpServer"
)

func Run() {
	// TODO: Write logging
	// TODO: Write repository
	// TODO: Write usecase
	// Init HTTP Server - controller
	// Init handler
	handler := httpHandler.New()
	// Init server
	httpServer.New(handler)
}
