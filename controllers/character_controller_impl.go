package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/trwalker/marvel-go/characters"
	"net/http"
)

var CharacterControllerInstance CharacterController = &CharacterControllerImpl{
	CharacterDetailsServiceInterface: characters.CharacterDetailsServiceInstance,
}

type CharacterControllerImpl struct {
	CharacterDetailsServiceInterface characters.CharacterDetailsService
}

func (controller *CharacterControllerImpl) Get(res http.ResponseWriter, req *http.Request) {
	routeVars := mux.Vars(req)

	characterName := routeVars["characterName"]

	characterDetailsModel, found, err := controller.CharacterDetailsServiceInterface.GetCharacter(characterName)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	} else if !found {
		res.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(res).Encode(characterDetailsModel)
	}
}