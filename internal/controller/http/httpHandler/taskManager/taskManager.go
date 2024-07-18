package taskManager

import (
	"github.com/sudo-odner/LifeManagerBack/internal/logger"
	"github.com/sudo-odner/LifeManagerBack/internal/usecase"
	"net/http"
)

type HandlerTaskManager struct {
	log       *logger.Logger
	u         *usecase.UseCase
	prefixURL string
}

func HandGroup(log *logger.Logger, u *usecase.UseCase, prefixURL string) http.Handler {
	h := &HandlerTaskManager{
		log:       log,
		u:         u,
		prefixURL: prefixURL,
	}

	mux := http.NewServeMux()

	return http.StripPrefix(h.prefixURL, mux)
}
