package characters

import (
	"sync"
)

var CharacterListServiceInstance CharacterListService = &CharacterListServiceImpl{
	CharacterMapRepoInterface: CharacterMapRepoInstance,
	CharacterServiceInterface: CharacterServiceInstance,
	lock: &sync.Mutex{},
	characterList: &CharacterListModel{
		Characters: make([]*CharacterModel, 0),
	},
}

type CharacterListServiceImpl struct {
	CharacterMapRepoInterface CharacterMapRepo
	CharacterServiceInterface CharacterService
	lock                      *sync.Mutex
	characterList             *CharacterListModel
}

type characterGetResult struct {
	Character *CharacterModel
	Found     bool
	Err       error
}

func (characterListService *CharacterListServiceImpl) GetCharacterList() *CharacterListModel {
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

func getCharacters(characterListService *CharacterListServiceImpl, characterMap map[string]int) []*CharacterModel {
	characterGetChannel := make(chan *characterGetResult)
	defer close(characterGetChannel)

	for name, _ := range characterMap {
		go getCharacter(characterListService, name, characterGetChannel)
	}

	var characters []*CharacterModel

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
