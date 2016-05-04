package charservices

import (
	"github.com/trwalker/marvel-go/characters/models"
	"github.com/trwalker/marvel-go/characters/repos"
	"sync"
)

var CharacterListServiceInstance CharacterListService = &CharacterListServiceImpl{
	CharacterMapRepoInterface: charrepos.CharacterMapRepoInstance,
	CharacterServiceInterface: CharacterServiceInstance,
	lock: &sync.Mutex{},
	characterList: &charmodels.CharacterListModel{
		Characters: make([]*charmodels.CharacterModel, 0),
	},
}

type CharacterListServiceImpl struct {
	CharacterMapRepoInterface charrepos.CharacterMapRepo
	CharacterServiceInterface CharacterService
	lock                      *sync.Mutex
	characterList             *charmodels.CharacterListModel
}

type characterGetResult struct {
	Character *charmodels.CharacterModel
	Found     bool
	Err       error
}

func (characterListService *CharacterListServiceImpl) GetCharacterList() *charmodels.CharacterListModel {
	if len(characterListService.characterList.Characters) == 0 {
		characterListService.lock.Lock()
		defer characterListService.lock.Unlock()

		if len(characterListService.characterList.Characters) == 0 {
			buildCharacterList(characterListService)
		}
	}

	return characterListService.characterList
}

func buildCharacterList(characterListService *CharacterListServiceImpl) {
	characterMap := characterListService.CharacterMapRepoInterface.GetCharacterMap()

	characters := getCharacters(characterListService, characterMap)

	characterListService.characterList.Characters = characters
}

func getCharacters(characterListService *CharacterListServiceImpl, characterMap map[string]int) []*charmodels.CharacterModel {
	characterGetChannel := make(chan *characterGetResult)
	defer close(characterGetChannel)

	for name, _ := range characterMap {
		go getCharacter(characterListService, name, characterGetChannel)
	}

	var characters []*charmodels.CharacterModel

	for i := 0; i < len(characterMap); i++ {
		result := <-characterGetChannel

		if result.Found && result.Err == nil {
			characters = append(characters, result.Character)
		}
	}

	return characters
}

func getCharacter(characterListService *CharacterListServiceImpl, name string, characterGetChannel chan *characterGetResult) {
	character, found, err := characterListService.CharacterServiceInterface.GetCharacter(name)

	result := &characterGetResult{
		Character: character,
		Found:     found,
		Err:       err,
	}

	characterGetChannel <- result
}
