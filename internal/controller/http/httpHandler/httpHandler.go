package httpHandler

import (
	"github.com/sudo-odner/LifeManagerBack/internal/controller/http/httpHandler/HandlerIMPL"
	"github.com/sudo-odner/LifeManagerBack/internal/logger"
	"github.com/sudo-odner/LifeManagerBack/internal/usecase"
	"net/http"
)

func Create(log *logger.Logger, u *usecase.UseCase) *http.ServeMux {
	h := handlerIMPL.New(log, u)

	mux := http.NewServeMux()
	// Check status work handler server
	mux.Handle("/health", http.HandlerFunc(h.Health))

	// SERVER
	// TODO: Write handler way
	// TODO: Write IMPL handler
	// TODO: Connect middleware
	return mux
}
