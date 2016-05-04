package charservices

import (
	"github.com/trwalker/marvel-go/auth"
	"github.com/trwalker/marvel-go/characters/models"
	"github.com/trwalker/marvel-go/characters/repos"
	"github.com/trwalker/marvel-go/rest"
)

var CharacterServiceInstance CharacterService = &CharacterServiceImpl{
	CharacterMapRepoInterface:   charrepos.CharacterMapRepoInstance,
	CharacterRepoInterface:      charrepos.CharacterRepoInstance,
	CharacterCacheRepoInterface: charrepos.CharacterCacheRepoInstance,
	CredentialsServiceInterface: auth.CredentialsServiceInstance,
	RestClientAdapterInterface:  rest.RestClientAdapterInstance,
}

type CharacterServiceImpl struct {
	CharacterMapRepoInterface   charrepos.CharacterMapRepo
	CharacterRepoInterface      charrepos.CharacterRepo
	CharacterCacheRepoInterface charrepos.CharacterCacheRepo
	CredentialsServiceInterface auth.CredentialsService
	RestClientAdapterInterface  rest.RestClientAdapter
}

func (characterService *CharacterServiceImpl) GetCharacter(name string) (character *charmodels.CharacterModel, found bool, err error) {
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
