package characters

import (
	"fmt"
	"github.com/trwalker/marvel-go/auth"
	"github.com/trwalker/marvel-go/rest"
	"time"
)

const getCharacterUrlFormat string = "http://gateway.marvel.com/v1/public/characters/%d?ts=%s&apikey=%s&hash=%s";
const timeout time.Duration = time.Millisecond * 4000

var CharacterDetailsServiceInstance CharacterDetailsService = &CharacterDetailsServiceImpl{
	CharacterMapRepoInterface:   CharacterMapRepoInstance,
	CredentialsServiceInterface: auth.CredentialsServiceInstance,
	RestClientAdapterInterface: rest.RestClientAdapterInstance,
}

type CharacterDetailsServiceImpl struct {
	CharacterMapRepoInterface   CharacterMapRepo
	CredentialsServiceInterface auth.CredentialsService
	RestClientAdapterInterface rest.RestClientAdapter
}

func (characterDetailsService *CharacterDetailsServiceImpl) GetCharacter(name string) *CharacterDetailsModel {
	// Get ID
	characterModel, found := TryGetCharacterByName(characterDetailsService, name)

	if found {
		// Get Credentials
		credentials := characterDetailsService.CredentialsServiceInterface.GenerateCredentials()

		// Build URL
		requestUrl := fmt.Sprintf(getCharacterUrlFormat, characterModel.Id, credentials.TimeStamp, credentials.PublicKey, credentials.Hash)

		// Make REST Call
		characterDetailsService.RestClientAdapterInterface.Get(requestUrl, timeout)

		// Parse response JSON into Model
	}

	return &CharacterDetailsModel{
		Name: characterModel.Name,
	}
}

func TryGetCharacterByName(characterDetailsService *CharacterDetailsServiceImpl, name string) (*CharacterModel, bool) {
	characterMap := characterDetailsService.CharacterMapRepoInterface.GetCharacterMap()
	characterModel, found := characterMap[name]

	return characterModel, found
}
