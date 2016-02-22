package config

import (
	"github.com/gorilla/mux"
	"net/http"
)

func BuildApiHandler() http.Handler {
	apiRouter := mux.NewRouter()

	RegisterRoutes(apiRouter)

	httpHandler := RegisterMiddleware(apiRouter)

	return httpHandler
}
