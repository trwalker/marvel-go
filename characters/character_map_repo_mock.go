package characters

var CharacterMapRepoMockInstance *CharacterMapRepoMock = &CharacterMapRepoMock{}

type CharacterMapRepoMock struct {
	characterMap map[string]int
}

func (characterMapRepoMock *CharacterMapRepoMock) ResetMock() {
	characterMapRepoMock.characterMap = make(map[string]int)
	characterMapRepoMock.characterMap["spider-man"] = 1
}

func (characterMapRepoMock *CharacterMapRepoMock) GetCharacterMapMockSetup(characterMapMock map[string]int) {
	characterMapRepoMock.characterMap = characterMapMock
}

func (characterMapRepoMock *CharacterMapRepoMock) GetCharacterMap() map[string]int {
	return characterMapRepoMock.characterMap
}
