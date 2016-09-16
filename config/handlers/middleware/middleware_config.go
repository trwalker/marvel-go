package middleware

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterMiddleware(apiRouter *mux.Router) http.Handler {
	var apiHandler http.Handler = apiRouter

	apiHandler = handlers.CompressHandler(apiHandler)
	apiHandler = handlers.CORS()(apiHandler)
	apiHandler = ResponseHeaders(apiHandler)

	return apiHandler
}
