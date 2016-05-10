package characterstest

import (
	"github.com/trwalker/marvel-go/characters"
)

type CharacterCacheRepoMock struct {
	GetMock func(id int) (character *characters.CharacterModel, found bool)
	AddMock func(character *characters.CharacterModel)
}

func (characterCacheRepoMock *CharacterCacheRepoMock) Get(id int) (character *characters.CharacterModel, found bool) {
	return characterCacheRepoMock.GetMock(id)
}

func (characterCacheRepoMock *CharacterCacheRepoMock) Add(character *characters.CharacterModel) {
	characterCacheRepoMock.AddMock(character)
}
