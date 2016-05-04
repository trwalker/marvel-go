package charrepos

import (
	"github.com/trwalker/marvel-go/auth"
	"github.com/trwalker/marvel-go/characters/models"
)

var CharacterRepoMockInstance *CharacterRepoMock = &CharacterRepoMock{}

type CharacterRepoMock struct {
	GetCharacterMock func(characterId int, credentials *auth.CredentialsModel) (character *charmodels.CharacterModel, found bool, err error)
}

func (characterRepoMock *CharacterRepoMock) GetCharacter(characterId int, credentials *auth.CredentialsModel) (character *charmodels.CharacterModel, found bool, err error) {
	return characterRepoMock.GetCharacterMock(characterId, credentials)
}
