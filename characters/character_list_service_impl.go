package characters

import (
	"sync"
	"strings"
)

var CharacterListServiceInstance CharacterListService = NewCharacterListService(CharacterServiceInstance, CharacterMapRepoInstance)

type characterListServiceImpl struct {
	characterServiceInterface CharacterService
	characterMapRepoInterface CharacterMapRepo
	characterList             *CharacterListModel
	lock                      *sync.Mutex
}

type characterGetResult struct {
	Character *CharacterModel
	Found     bool
	Err       error
}

func NewCharacterListService(characterService CharacterService, characterMapRepo CharacterMapRepo) CharacterListService {
	characterListService := &characterListServiceImpl{
		characterServiceInterface: characterService,
		characterMapRepoInterface: characterMapRepo,
		characterList:             &CharacterListModel{Characters: make([]*CharacterModel, 0)},
		lock:                      &sync.Mutex{},
	}

	return characterListService
}

func (characterListService *characterListServiceImpl) GetCharacterList(filter string) *CharacterListModel {
	if len(characterListService.characterList.Characters) == 0 {
		characterListService.lock.Lock()
		defer characterListService.lock.Unlock()

		if len(characterListService.characterList.Characters) == 0 {
			buildCharacterList(characterListService)
		}
	}

	return getFilteredCharacterList(characterListService.characterList, filter)
}

func buildCharacterList(characterListService *characterListServiceImpl) {
	characterMap := characterListService.characterMapRepoInterface.GetCharacterMap()

	characters := getCharacters(characterListService, characterMap)

	characterListService.characterList.Characters = characters
}

func getCharacters(characterListService *characterListServiceImpl, characterMap map[string]int) []*CharacterModel {
	characterGetChannel := make(chan *characterGetResult)
	defer close(characterGetChannel)

	for name := range characterMap {
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

func getFilteredCharacterList(characterList *CharacterListModel, filter string) *CharacterListModel {
	if filter == "" {
		return characterList
	} else {
		filteredCharacterList := &CharacterListModel{Characters: make([]*CharacterModel, 0)}

		filterLower := strings.ToLower(filter)

		for _, character := range characterList.Characters {
			if strings.Contains(strings.ToLower(character.Name), filterLower) {
				filteredCharacterList.Characters = append(filteredCharacterList.Characters, character)
			}
		}

		return filteredCharacterList
	}
}
