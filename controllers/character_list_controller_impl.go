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
		characterListServiceInterface: characterListService,
	}

	return characterListController
}

func (controller *characterListControllerImpl) Get(res http.ResponseWriter, req *http.Request) {
	var filter string

	filterQuery := req.URL.Query()["filter"]
	if len(filterQuery) == 0 {
		filter = ""
	} else {
		filter = filterQuery[0]
	}

	characterListModel := controller.characterListServiceInterface.GetCharacterList(filter)

	if characterListModel != nil {
		json.NewEncoder(res).Encode(characterListModel)
	} else {
		http.NotFound(res, req)
	}
}
