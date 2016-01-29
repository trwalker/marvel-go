package controllers

import (
	"encoding/json"
	"github.com/trwalker/marvel-go/services"
	"net/http"
)

type CharacterListController struct {
	CharacterListServiceInterface services.CharacterListService
}

func Constructor() *CharacterListController {
	return &CharacterListController{
		CharacterListServiceInterface: services.Constructor(),
	}
}

func (controller *CharacterListController) Get(res http.ResponseWriter, req *http.Request) {
	characterListModel := controller.CharacterListServiceInterface.GetCharacterList()

	json.NewEncoder(res).Encode(characterListModel)
}
