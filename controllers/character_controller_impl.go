package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/trwalker/marvel-go/characters/services"
	"net/http"
)

var CharacterControllerInstance CharacterController = &CharacterControllerImpl{
	CharacterServiceInterface: charservices.CharacterServiceInstance,
}

type CharacterControllerImpl struct {
	CharacterServiceInterface charservices.CharacterService
}

func (controller *CharacterControllerImpl) Get(res http.ResponseWriter, req *http.Request) {
	routeVars := mux.Vars(req)

	characterName := routeVars["characterName"]

	characterModel, found, err := controller.CharacterServiceInterface.GetCharacter(characterName)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	} else if !found {
		res.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(res).Encode(characterModel)
	}
}
