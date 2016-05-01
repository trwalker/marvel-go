package characters

var CharacterServiceMockInstance *CharacterServiceMock = &CharacterServiceMock{}

type CharacterServiceMock struct {
	character *CharacterModel
	found     bool
	err       error
}

func (characterServiceMock *CharacterServiceMock) ResetMock() {
	characterServiceMock.character = &CharacterModel{
	 	Id:          1,
		Name:        "spider-man",
		Description: "amazing spider man",
		Image:       "https://cdn.com/spidey.jpg",
	}
	characterServiceMock.found = true
	characterServiceMock.err = nil
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
