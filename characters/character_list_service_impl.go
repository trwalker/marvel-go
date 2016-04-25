package characters

import (
	"sync"
)

var CharacterListServiceInstance CharacterListService = &CharacterListServiceImpl{
	CharacterMapRepoInterface: CharacterMapRepoInstance,
	CharacterServiceInterface: CharacterServiceInstance,

	characterList: &CharacterListModel{
		Characters: make([]*CharacterModel, 0),
	},
}

type CharacterListServiceImpl struct {
	CharacterMapRepoInterface CharacterMapRepo
	CharacterServiceInterface CharacterService
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

	for name, _ := range characterMap {
		character, found, err := characterListService.CharacterServiceInterface.GetCharacter(name)

		if found && err == nil {
			characters = append(characters, character)
		}
	}

	characterListService.characterList.Characters = characters
}
