package services

import (
	"github.com/trwalker/marvel-go/models"
	"github.com/trwalker/marvel-go/repos"
	"sync"
)

var CharacterListServiceInstance CharacterListService = constructor()

var characterList models.CharacterListModel = models.CharacterListModel{}

type CharacterListServiceImpl struct {
	CharacterMapRepoInterface repos.CharacterMapRepo
}

func constructor() CharacterListService {
	var characterListService *CharacterListServiceImpl = &CharacterListServiceImpl{
		CharacterMapRepoInterface: repos.CharacterMapRepoInstance,
	}

	return characterListService
}

func (characterListService *CharacterListServiceImpl) GetCharacterList() models.CharacterListModel {
	if len(characterList.Characters) == 0 {
		lock := &sync.Mutex{}

		lock.Lock()
		defer lock.Unlock()

		if len(characterList.Characters) == 0 {
			buildCharacterList(characterListService)
		}
	}

	return characterList
}

func buildCharacterList(characterListService *CharacterListServiceImpl) {
	characterMap := characterListService.CharacterMapRepoInterface.GetCharacterMap()

	var characters []*models.CharacterModel

	for _, value := range characterMap {
		characters = append(characters, value)
	}

	characterList.Characters = characters
}
