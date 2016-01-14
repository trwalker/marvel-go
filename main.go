package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/trwalker/marvel-go/controllers"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	characterListController := controllers.CharacterListController{}
	router.HandleFunc("/v1/characters", characterListController.Get).Methods("GET")

	characterController := controllers.CharacterController{}
	router.HandleFunc("/v1/characters/{characterName}", characterController.Get).Methods("GET")

	router.Headers("Content-Type", "application/json")

	handlers.CompressHandler(router)

	fmt.Println("Starting Web Server...")
	fmt.Println("URL:", "http://127.0.0.1:9000/")
	http.ListenAndServe("127.0.0.1:9000", handlers.CORS()(router))
}
