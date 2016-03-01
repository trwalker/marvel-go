package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/trwalker/marvel-go/characters"
	"net/http"
)

var CharacterControllerInstance CharacterController = &CharacterControllerImpl{}

type CharacterControllerImpl struct {
}

func (controller *CharacterControllerImpl) Get(res http.ResponseWriter, req *http.Request) {
	routeVars := mux.Vars(req)

	model := &characters.CharacterDetailsModel{Name: routeVars["characterName"], Id: 1}

	json.NewEncoder(res).Encode(model)
}
