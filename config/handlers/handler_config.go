package handlers

import (
	"github.com/gorilla/mux"
	"github.com/trwalker/marvel-go/config/handlers/middleware"
	"github.com/trwalker/marvel-go/config/handlers/routing"
	"net/http"
)

func BuildApiHandler() http.Handler {
	apiRouter := mux.NewRouter()

	routing.RegisterRoutes(apiRouter)

	apiHandler := middleware.RegisterMiddleware(apiRouter)

	return apiHandler
}
