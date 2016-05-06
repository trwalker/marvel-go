package characterstests

import (
	"github.com/trwalker/marvel-go/characters"
)

var CharacterListServiceMockInstance *CharacterListServiceMock = &CharacterListServiceMock{}

type CharacterListServiceMock struct {
	GetCharacterListMock func() *characters.CharacterListModel
}

func (characterListServiceMock *CharacterListServiceMock) GetCharacterList() *characters.CharacterListModel {
	return characterListServiceMock.GetCharacterListMock()
}
