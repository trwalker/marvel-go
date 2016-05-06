package characters

import (
	"github.com/trwalker/marvel-go/auth"
)

var CharacterServiceInstance CharacterService = NewCharacterService(CharacterMapRepoInstance, CharacterRepoInstance, CharacterCacheRepoInstance, auth.CredentialsServiceInstance)

type characterServiceImpl struct {
	characterMapRepoInterface   CharacterMapRepo
	characterRepoInterface      CharacterRepo
	characterCacheRepoInterface CharacterCacheRepo
	credentialsServiceInterface auth.CredentialsService
}

func NewCharacterService(
	characterMapRepo CharacterMapRepo,
	characterRepo CharacterRepo,
	characterCacheRepo CharacterCacheRepo,
	credentialsService auth.CredentialsService) CharacterService {

	characterService := &characterServiceImpl{
		characterMapRepoInterface:   characterMapRepo,
		characterRepoInterface:      characterRepo,
		characterCacheRepoInterface: characterCacheRepo,
		credentialsServiceInterface: credentialsService,
	}

	return characterService
}

func (characterService *characterServiceImpl) GetCharacter(name string) (character *CharacterModel, found bool, err error) {
	character = nil
	found = false
	err = nil

	characterId, mappingFound := tryGetCharacterByName(characterService, name)

	if mappingFound {
		character, found = characterService.characterCacheRepoInterface.Get(characterId)

		if !found || character == nil {
			credentials := characterService.credentialsServiceInterface.GenerateCredentials()
			character, found, err = characterService.characterRepoInterface.GetCharacter(characterId, credentials)

			if found && character != nil {
				characterService.characterCacheRepoInterface.Add(character)
			}
		}
	}

	return
}

func tryGetCharacterByName(characterService *characterServiceImpl, name string) (characterId int, found bool) {
	characterMap := characterService.characterMapRepoInterface.GetCharacterMap()
	characterId, found = characterMap[name]

	return
}
