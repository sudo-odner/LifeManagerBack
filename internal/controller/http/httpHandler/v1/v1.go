package v1

import "net/http"

func Group() http.Handler {
	mux := http.NewServeMux()

	return http.StripPrefix("/v1", mux)
}
