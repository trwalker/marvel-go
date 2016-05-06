package controllers

import (
	"encoding/json"
	"github.com/trwalker/marvel-go/characters"
	"net/http"
)

var CharacterListControllerInstance CharacterListController = NewCharacterListController(characters.CharacterListServiceInstance)

type characterListControllerImpl struct {
	characterListServiceInterface characters.CharacterListService
}

func NewCharacterListController(characterListService characters.CharacterListService) CharacterListController {
	characterListController := &characterListControllerImpl{
		characterListServiceInterface: characters.CharacterListServiceInstance,
	}

	return characterListController
}

func (controller *characterListControllerImpl) Get(res http.ResponseWriter, req *http.Request) {
	characterListModel := controller.characterListServiceInterface.GetCharacterList()

	if characterListModel != nil {
		json.NewEncoder(res).Encode(characterListModel)
	} else {
		http.NotFound(res, req)
	}
}
