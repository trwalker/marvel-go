package repos

import (
	"github.com/trwalker/marvel-go/models"
	"testing"
)

type CharMapRepoTestContext struct {
	Repo CharacterMapRepo
}

var charMapRepoTestContext *CharMapRepoTestContext = new(CharMapRepoTestContext)

func (context *CharMapRepoTestContext) Setup() {
	context.Repo = Constructor()
}

func (context *CharMapRepoTestContext) TearDown() {
	context = nil
}

func TestCharacterMapRepoValidState(t *testing.T) {
	charMapRepoTestContext.Setup()
	defer charMapRepoTestContext.TearDown()

	characterMap := charMapRepoTestContext.Repo.GetCharacterMap()

	assertCharacterMapLength(t, characterMap)
	assertCharacterMapValue(t, characterMap)
}

func assertCharacterMapLength(t *testing.T, characterMap map[string]models.CharacterModel) {
	if len(characterMap) != 11 {
		t.Errorf("Character map should contain 11 characters, length: %v", len(characterMap))
	}
}

func assertCharacterMapValue(t *testing.T, characterMap map[string]models.CharacterModel) {
	_, found := characterMap["spider-man"]

	if !found {
		t.Error("Character map should contain \"spider-man\"")
	}
}
