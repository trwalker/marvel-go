package characterstests

import (
	"github.com/trwalker/marvel-go/auth"
	"github.com/trwalker/marvel-go/characters"
)

var CharacterRepoMockInstance *CharacterRepoMock = &CharacterRepoMock{}

type CharacterRepoMock struct {
	GetCharacterMock func(characterId int, credentials *auth.CredentialsModel) (character *characters.CharacterModel, found bool, err error)
}

func (characterRepoMock *CharacterRepoMock) GetCharacter(characterId int, credentials *auth.CredentialsModel) (character *characters.CharacterModel, found bool, err error) {
	return characterRepoMock.GetCharacterMock(characterId, credentials)
}
