package characterstests

var CharacterMapRepoMockInstance *CharacterMapRepoMock = &CharacterMapRepoMock{}

type CharacterMapRepoMock struct {
	GetCharacterMapMock func() map[string]int
}

func (characterMapRepoMock *CharacterMapRepoMock) GetCharacterMap() map[string]int {
	return characterMapRepoMock.GetCharacterMapMock()
}
