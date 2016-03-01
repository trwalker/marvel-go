package controllers

import (
	"encoding/json"
	"github.com/trwalker/marvel-go/characters"
	"net/http"
)

var CharacterListControllerInstance CharacterListController = &CharacterListControllerImpl{
	CharacterListServiceInterface: characters.CharacterListServiceInstance,
}

type CharacterListControllerImpl struct {
	CharacterListServiceInterface characters.CharacterListService
}

func (controller *CharacterListControllerImpl) Get(res http.ResponseWriter, req *http.Request) {
	characterListModel := controller.CharacterListServiceInterface.GetCharacterList()

	if characterListModel != nil {
		json.NewEncoder(res).Encode(characterListModel)
	} else {
		http.NotFound(res, req)
	}
}
