package controllers

import (
	"encoding/json"
	"github.com/trwalker/marvel-go/characters/services"
	"net/http"
)

var CharacterListControllerInstance CharacterListController = &CharacterListControllerImpl{
	CharacterListServiceInterface: charservices.CharacterListServiceInstance,
}

type CharacterListControllerImpl struct {
	CharacterListServiceInterface charservices.CharacterListService
}

func (controller *CharacterListControllerImpl) Get(res http.ResponseWriter, req *http.Request) {
	characterListModel := controller.CharacterListServiceInterface.GetCharacterList()

	if characterListModel != nil {
		json.NewEncoder(res).Encode(characterListModel)
	} else {
		http.NotFound(res, req)
	}
}
