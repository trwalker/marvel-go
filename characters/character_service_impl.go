package characters

import (
	"github.com/trwalker/marvel-go/auth"
	"github.com/trwalker/marvel-go/rest"
	"time"
)

const getCharacterUrlFormat string = "http://gateway.marvel.com/v1/public/characters/%d?ts=%s&apikey=%s&hash=%s"
const timeout time.Duration = time.Millisecond * 4000

var CharacterServiceInstance CharacterService = &CharacterServiceImpl{
	CharacterMapRepoInterface:   CharacterMapRepoInstance,
	CharacterRepoInterface:      CharacterRepoInstance,
	CharacterCacheRepoInterface: CharacterCacheRepoInstance,
	CredentialsServiceInterface: auth.CredentialsServiceInstance,
	RestClientAdapterInterface:  rest.RestClientAdapterInstance,
}

type CharacterServiceImpl struct {
	CharacterMapRepoInterface   CharacterMapRepo
	CharacterRepoInterface      CharacterRepo
	CharacterCacheRepoInterface CharacterCacheRepo
	CredentialsServiceInterface auth.CredentialsService
	RestClientAdapterInterface  rest.RestClientAdapter
	characterCache              map[int]*CharacterModel
}

func (characterService *CharacterServiceImpl) GetCharacter(name string) (character *CharacterModel, found bool, err error) {
	character = nil
	found = false
	err = nil

	characterId, mappingFound := tryGetCharacterByName(characterService, name)

	if mappingFound {
		character, found = characterService.CharacterCacheRepoInterface.Get(characterId)

		if !found || character == nil {
			credentials := characterService.CredentialsServiceInterface.GenerateCredentials()
			character, found, err = characterService.CharacterRepoInterface.GetCharacter(characterId, credentials)

			if found && character != nil {
				characterService.CharacterCacheRepoInterface.Add(character)
			}
		}
	}

	return
}

func tryGetCharacterByName(characterService *CharacterServiceImpl, name string) (characterId int, found bool) {
	characterMap := characterService.CharacterMapRepoInterface.GetCharacterMap()
	characterId, found = characterMap[name]

	return
}
