package httpHandler

import (
	"github.com/sudo-odner/LifeManagerBack/internal/controller/http/httpHandler/v1"
	"github.com/sudo-odner/LifeManagerBack/internal/logger"
	"github.com/sudo-odner/LifeManagerBack/internal/usecase"
	"net/http"
)

type Handler struct {
	log logger.Logger
	u   usecase.UseCase
}

func New() *http.ServeMux {
	mux := http.NewServeMux()
	// Check status work handler server
	// todo: mux.Handle("/health", func)
	// todo: Authorization
	// Version server
	// TODO: create middleware
	mux.Handle("/v1/", v1.Group())

	// TODO: create middleware
	return mux
}
