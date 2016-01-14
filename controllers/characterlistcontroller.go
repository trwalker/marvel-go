package controllers

import (
	"encoding/json"
	"github.com/trwalker/marvel-go/models"
	"net/http"
)

type CharacterListController struct {
}

func (controller *CharacterListController) Get(res http.ResponseWriter, req *http.Request) {
	characterModel := &models.CharacterModel{Name: "spider-man", Id: 1}

	characterListModel := &models.CharacterListModel{}
	characterListModel.Characters = append(characterListModel.Characters, characterModel)

	json.NewEncoder(res).Encode(characterListModel)
}
