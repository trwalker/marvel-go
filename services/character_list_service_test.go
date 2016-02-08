package services

import (
	"github.com/trwalker/marvel-go/models"
	"testing"
)

var charListServiceTestContext *CharListServiceTestContext = new(CharListServiceTestContext)

type CharListServiceTestContext struct {
	Service      CharacterListService
	CharacterMap map[string]*models.CharacterModel
}

type CharacterMapRepoMock struct {
}

func (charMapRepoMock *CharacterMapRepoMock) GetCharacterMap() map[string]*models.CharacterModel {
	return charListServiceTestContext.CharacterMap
}

func (context *CharListServiceTestContext) Setup() {
	charListServiceTestContext.CharacterMap = make(map[string]*models.CharacterModel)

	charListServiceTestContext.CharacterMap["spider-man"] = &models.CharacterModel{
		Id:    1,
		Name:  "spider-man",
		Image: "http://i.annihil.us/u/prod/marvel/bar.jpg",
	}

	charListServiceTestContext.CharacterMap["hulk"] = &models.CharacterModel{
		Id:    2,
		Name:  "hulk",
		Image: "http://i.annihil.us/u/prod/marvel/foo.jpg",
	}

	charListServiceTestContext.Service = &CharacterListServiceImpl{
		CharacterMapRepoInterface: &CharacterMapRepoMock{},
	}
}

func (context *CharListServiceTestContext) TearDown() {
}

func TestCharacterListServiceInvalidState(t *testing.T) {
	charListServiceTestContext.Setup()
	defer charListServiceTestContext.TearDown()

	charListServiceTestContext.CharacterMap = nil

	characterList := charListServiceTestContext.Service.GetCharacterList()

	assertCharacterListIsEmpty(t, characterList)
}

func TestCharacterListServiceValidState(t *testing.T) {
	charListServiceTestContext.Setup()
	defer charListServiceTestContext.TearDown()

	characterList := charListServiceTestContext.Service.GetCharacterList()

	assertCharacterListLength(t, characterList)
	assertCharacterListContainsMapCharacters(t, characterList)
}

func assertCharacterListIsEmpty(t *testing.T, characterList models.CharacterListModel) {
	if len(characterList.Characters) != 0 {
		t.Error("Character list should be empty")
	}
}

func assertCharacterListLength(t *testing.T, characterList models.CharacterListModel) {
	if len(characterList.Characters) != len(charListServiceTestContext.CharacterMap) {
		t.Error("Character list length should match character map length")
	}
}

func assertCharacterListContainsMapCharacters(t *testing.T, characterList models.CharacterListModel) {
	for _, listChar := range characterList.Characters {
		_, found := charListServiceTestContext.CharacterMap[listChar.Name]
		if !found {
			t.Error("List character not found in map: ", listChar.Name)
		}
	}
}
