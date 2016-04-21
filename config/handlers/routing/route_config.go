package routing

import (
	"github.com/gorilla/mux"
	"github.com/trwalker/marvel-go/controllers"
	"net/http"
)

func RegisterRoutes(apiRouter *mux.Router) {
	apiRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	apiRouter.HandleFunc("/v1/characters", controllers.CharacterListControllerInstance.Get).Methods("GET")
	apiRouter.HandleFunc("/v1/characters/{characterName}", controllers.CharacterControllerInstance.Get).Methods("GET")
}
