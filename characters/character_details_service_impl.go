package characters

import (
	"github.com/trwalker/marvel-go/auth"
)

var CharacterDetailsServiceInstance CharacterDetailsService = &CharacterDetailsServiceImpl{
	CharacterMapRepoInterface:   CharacterMapRepoInstance,
	CredentialsServiceInterface: auth.CredentialsServiceInstance,
}

type CharacterDetailsServiceImpl struct {
	CharacterMapRepoInterface   CharacterMapRepo
	CredentialsServiceInterface auth.CredentialsService
}

func (characterDetailsService *CharacterDetailsServiceImpl) GetCharacter(name string) *CharacterDetailsModel {
	// Get ID
	characterModel, found := TryGetCharacterByName(characterDetailsService, name)

	if found {
		// Get Credentials
		// Make REST Call
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
