package controllers

import (
	"encoding/json"
	"github.com/trwalker/marvel-go/services"
	"net/http"
)

var CharacterListControllerInstance *CharacterListController = constructor()

type CharacterListController struct {
	CharacterListServiceInterface services.CharacterListService
}

func constructor() *CharacterListController {
	return &CharacterListController{
		CharacterListServiceInterface: services.CharacterListServiceInstance,
	}
}

func (controller *CharacterListController) Get(res http.ResponseWriter, req *http.Request) {
	characterListModel := controller.CharacterListServiceInterface.GetCharacterList()

	if characterListModel != nil {
		json.NewEncoder(res).Encode(characterListModel)
	} else {
		http.NotFound(res, req)
	}
}
