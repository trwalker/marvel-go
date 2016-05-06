package characters

import (
	"sync"
)

var CharacterListServiceInstance CharacterListService = NewCharacterListService(CharacterMapRepoInstance, CharacterServiceInstance)

type characterListServiceImpl struct {
	characterMapRepoInterface CharacterMapRepo
	characterServiceInterface CharacterService
	characterList             *CharacterListModel
	lock                      *sync.Mutex
}

type characterGetResult struct {
	Character *CharacterModel
	Found     bool
	Err       error
}

func NewCharacterListService(characterMapRepo CharacterMapRepo, characterService CharacterService) CharacterListService {
	characterListService := &characterListServiceImpl{
		characterMapRepoInterface: characterMapRepo,
		characterServiceInterface: characterService,
		characterList:             &CharacterListModel{Characters: make([]*CharacterModel, 0)},
		lock:                      &sync.Mutex{},
	}

	return characterListService
}

func (characterListService *characterListServiceImpl) GetCharacterList() *CharacterListModel {
	if len(characterListService.characterList.Characters) == 0 {
		characterListService.lock.Lock()
		defer characterListService.lock.Unlock()

		if len(characterListService.characterList.Characters) == 0 {
			buildCharacterList(characterListService)
		}
	}

	return characterListService.characterList
}

func buildCharacterList(characterListService *characterListServiceImpl) {
	characterMap := characterListService.characterMapRepoInterface.GetCharacterMap()

	characters := getCharacters(characterListService, characterMap)

	characterListService.characterList.Characters = characters
}

func getCharacters(characterListService *characterListServiceImpl, characterMap map[string]int) []*CharacterModel {
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

func getCharacter(characterListService *characterListServiceImpl, name string, characterGetChannel chan *characterGetResult) {
	character, found, err := characterListService.characterServiceInterface.GetCharacter(name)

	result := &characterGetResult{
		Character: character,
		Found:     found,
		Err:       err,
	}

	characterGetChannel <- result
}
