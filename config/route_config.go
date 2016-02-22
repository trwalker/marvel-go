package config

import (
	"github.com/gorilla/mux"
	"github.com/trwalker/marvel-go/controllers"
)

func RegisterRoutes(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/v1/characters", controllers.CharacterListControllerInstance.Get).Methods("GET")
	apiRouter.HandleFunc("/v1/characters/{characterName}", controllers.CharacterControllerInstance.Get).Methods("GET")
}
