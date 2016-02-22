package config

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/trwalker/marvel-go/middleware"
	"net/http"
)

func registerMiddleware(apiRouter *mux.Router) http.Handler {
	var apiHandler http.Handler = apiRouter

	apiHandler = handlers.CompressHandler(apiHandler)
	apiHandler = handlers.CORS(handlers.AllowedOrigins([]string{"http://google.com"}))(apiHandler)
	apiHandler = middleware.ResponseHeaders(apiHandler)

	return apiHandler
}
