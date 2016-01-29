package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/trwalker/marvel-go/controllers"
	"github.com/trwalker/marvel-go/middleware"
	"net/http"
)

func main() {
	apiHandler := initializeApiHandler()

	startServer(apiHandler)
}

func initializeApiHandler() http.Handler {
	apiRouter := mux.NewRouter()

	registerRoutes(apiRouter)

	apiHandler := registerMiddleware(apiRouter)

	return apiHandler
}

func registerRoutes(apiRouter *mux.Router) {
	characterListController := controllers.Constructor()
	apiRouter.HandleFunc("/v1/characters", characterListController.Get).Methods("GET")

	characterController := controllers.CharacterController{}
	apiRouter.HandleFunc("/v1/characters/{characterName}", characterController.Get).Methods("GET")
}

func registerMiddleware(apiRouter *mux.Router) http.Handler {
	var apiHandler http.Handler = apiRouter

	apiHandler = handlers.CompressHandler(apiHandler)
	apiHandler = handlers.CORS(handlers.AllowedOrigins([]string{"http://google.com"}))(apiHandler)
	apiHandler = middleware.ResponseHeaders(apiHandler)

	return apiHandler
}

func startServer(apiHandler http.Handler) {
	fmt.Println("Starting Web Server...")
	fmt.Println("URL:", "http://127.0.0.1:9000/")

	http.ListenAndServe("127.0.0.1:9000", apiHandler)
}
