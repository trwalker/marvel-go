package services

import (
	"github.com/trwalker/marvel-go/models"
	"github.com/trwalker/marvel-go/repos"
	"sync"
)

var CharacterListServiceInstance CharacterListService = &CharacterListServiceImpl{
	CharacterMapRepoInterface: repos.CharacterMapRepoInstance,
	characterList: &models.CharacterListModel{
		Characters: make([]*models.CharacterModel, 0),
	},
}

type CharacterListServiceImpl struct {
	CharacterMapRepoInterface repos.CharacterMapRepo
	characterList             *models.CharacterListModel
}

func (characterListService *CharacterListServiceImpl) GetCharacterList() *models.CharacterListModel {
	if len(characterListService.characterList.Characters) == 0 {
		lock := &sync.Mutex{}

		lock.Lock()
		defer lock.Unlock()

		if len(characterListService.characterList.Characters) == 0 {
			buildCharacterList(characterListService)
		}
	}

	return characterListService.characterList
}

func buildCharacterList(characterListService *CharacterListServiceImpl) {
	characterMap := characterListService.CharacterMapRepoInterface.GetCharacterMap()

	var characters []*models.CharacterModel

	for _, value := range characterMap {
		characters = append(characters, value)
	}

	characterListService.characterList.Characters = characters
}
