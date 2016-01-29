package services

import (
	"github.com/trwalker/marvel-go/models"
	"github.com/trwalker/marvel-go/repos"
)

type CharacterListServiceImpl struct {
	CharacterMapRepoInterface repos.CharacterMapRepo
}

func Constructor() CharacterListService {
	var characterListService *CharacterListServiceImpl = &CharacterListServiceImpl{
		CharacterMapRepoInterface: repos.Constructor(),
	}

	return characterListService
}

func (characterListService *CharacterListServiceImpl) GetCharacterList() models.CharacterListModel {
	characterMap := characterListService.CharacterMapRepoInterface.GetCharacterMap()

	var characterList []*models.CharacterModel

	for _, value := range characterMap {
		characterList = append(characterList, value)
	}

	return models.CharacterListModel{Characters: characterList}
}
