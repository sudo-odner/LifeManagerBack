package httpHandler

import (
	"github.com/sudo-odner/LifeManagerBack/internal/controller/http/httpHandler/HandlerIMPL"
	"github.com/sudo-odner/LifeManagerBack/internal/controller/http/httpHandler/taskManager"
	"github.com/sudo-odner/LifeManagerBack/internal/logger"
	"github.com/sudo-odner/LifeManagerBack/internal/usecase"
	"net/http"
)

func Create(log *logger.Logger, u *usecase.UseCase) *http.ServeMux {
	h := handlerIMPL.New(log, u)

	mux := http.NewServeMux()
	// Check status work handler server
	// TODO: create middleware
	// todo: mux.Handle("/health", func)
	// todo: Authorization
	mux.Handle("/health", http.HandlerFunc(h.Health))

	// SERVER
	// Handler Task manager
	handlerTaskManager := taskManager.HandGroup(log, u, "/v1")
	mux.Handle("/v1/", handlerTaskManager)
	return mux
}
