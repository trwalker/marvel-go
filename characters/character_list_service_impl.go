package characters

import (
	"sync"
)

var CharacterListServiceInstance CharacterListService = &CharacterListServiceImpl{
	CharacterMapRepoInterface: CharacterMapRepoInstance,
	characterList: &CharacterListModel{
		Characters: make([]*CharacterModel, 0),
	},
}

type CharacterListServiceImpl struct {
	CharacterMapRepoInterface CharacterMapRepo
	characterList             *CharacterListModel
}

func (characterListService *CharacterListServiceImpl) GetCharacterList() *CharacterListModel {
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

	var characters []*CharacterModel

	for _, value := range characterMap {
		characters = append(characters, value)
	}

	characterListService.characterList.Characters = characters
}
