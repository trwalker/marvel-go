package charservices

import (
	"github.com/trwalker/marvel-go/characters/models"
)

var CharacterServiceMockInstance *CharacterServiceMock = &CharacterServiceMock{}

type CharacterServiceMock struct {
	GetCharacterMock func(name string) (character *charmodels.CharacterModel, found bool, err error)
}

func (characterServiceMock *CharacterServiceMock) GetCharacter(name string) (character *charmodels.CharacterModel, found bool, err error) {
	return characterServiceMock.GetCharacterMock(name)
}
