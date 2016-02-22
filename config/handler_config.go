package config

import (
	"github.com/gorilla/mux"
	"net/http"
)

func BuildApiHandler() http.Handler {
	apiRouter := mux.NewRouter()

	registerRoutes(apiRouter)

	httpHandler := registerMiddleware(apiRouter)

	return httpHandler
}
