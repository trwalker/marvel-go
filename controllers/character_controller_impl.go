package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/trwalker/marvel-go/characters"
	"net/http"
)

var CharacterControllerInstance CharacterController = NewCharacterController(characters.CharacterServiceInstance)

type characterControllerImpl struct {
	characterServiceInterface characters.CharacterService
}

func NewCharacterController(characterService characters.CharacterService) CharacterController {
	characterController := &characterControllerImpl{
		characterServiceInterface: characterService,
	}

	return characterController
}

func (controller *characterControllerImpl) Get(res http.ResponseWriter, req *http.Request) {
	routeVars := mux.Vars(req)

	characterName := routeVars["characterName"]

	characterModel, found, err := controller.characterServiceInterface.GetCharacter(characterName)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	} else if !found {
		res.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(res).Encode(characterModel)
	}
}
