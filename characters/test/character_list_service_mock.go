package characterstest

import (
	"github.com/trwalker/marvel-go/characters"
)

type CharacterListServiceMock struct {
	GetCharacterListMock func() *characters.CharacterListModel
}

func (characterListServiceMock *CharacterListServiceMock) GetCharacterList() *characters.CharacterListModel {
	return characterListServiceMock.GetCharacterListMock()
}
