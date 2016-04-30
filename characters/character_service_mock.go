package characters

var CharacterServiceMockInstance *CharacterServiceMock = &CharacterServiceMock{}

type CharacterServiceMock struct {
	character *CharacterModel
	found     bool
	err       error
}

func (characterServiceMock *CharacterServiceMock) GetCharacterMockSetup(characterMock *CharacterModel, foundMock bool, errMock error) {
	characterServiceMock.character = characterMock
	characterServiceMock.found = foundMock
	characterServiceMock.err = errMock
}

func (characterServiceMock *CharacterServiceMock) GetCharacter(name string) (character *CharacterModel, found bool, err error) {
	character = characterServiceMock.character
	found = characterServiceMock.found
	err = characterServiceMock.err

	return
}
