package characterstests

import (
	"github.com/trwalker/marvel-go/characters"
)

var CharacterServiceMockInstance *CharacterServiceMock = &CharacterServiceMock{}

type CharacterServiceMock struct {
	GetCharacterMock func(name string) (character *characters.CharacterModel, found bool, err error)
}

func (characterServiceMock *CharacterServiceMock) GetCharacter(name string) (character *characters.CharacterModel, found bool, err error) {
	return characterServiceMock.GetCharacterMock(name)
}
