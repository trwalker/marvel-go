package repos

import (
	"github.com/trwalker/marvel-go/models"
	"testing"
)

var charMapRepoTestContext *CharMapRepoTestContext = new(CharMapRepoTestContext)

type CharMapRepoTestContext struct {
	Repo CharacterMapRepo
}

func (context *CharMapRepoTestContext) Setup() {
	context.Repo = CharacterMapRepoInstance
}

func (context *CharMapRepoTestContext) TearDown() {
}

func TestCharacterMapRepoWhenValidState(t *testing.T) {
	charMapRepoTestContext.Setup()
	defer charMapRepoTestContext.TearDown()

	characterMap := charMapRepoTestContext.Repo.GetCharacterMap()

	assertCharacterMapLength(t, characterMap)
	assertCharacterMapValue(t, characterMap)
}

func assertCharacterMapLength(t *testing.T, characterMap map[string]*models.CharacterModel) {
	if len(characterMap) != 11 {
		t.Errorf("Character map should contain 11 characters, length: %v", len(characterMap))
	}
}

func assertCharacterMapValue(t *testing.T, characterMap map[string]*models.CharacterModel) {
	_, found := characterMap["spider-man"]

	if !found {
		t.Error("Character map should contain \"spider-man\"")
	}
}
