package characterstest

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/trwalker/marvel-go/characters"
	"testing"
)

func TestCharacterListServiceSpec(t *testing.T) {

	Convey("CharacterListService Tests", t, func() {

		characterMap := make(map[string]*characters.CharacterModel)
		characterMap["spider-man"] = &characters.CharacterModel{
			Id:          1,
			Name:        "spider-man",
			Description: "amazing spider man",
			Image:       "https://cdn.com/spidey.jpg",
		}

		characterMap["hulk"] = &characters.CharacterModel{
			Id:          2,
			Name:        "hulk",
			Description: "don't make me angry",
			Image:       "https://cdn.com/hulk.jpg",
		}

		var getCharacterError error = nil

		characterServiceMock := &CharacterServiceMock{}
		characterServiceMock.GetCharacterMock = func(name string) (character *characters.CharacterModel, found bool, err error) {
			character, found = characterMap[name]
			err = getCharacterError

			return
		}

		characterMapRepoMock := &CharacterMapRepoMock{}
		characterMapRepoMock.GetCharacterMapMock = func() map[string]int {
			characterIdMap := make(map[string]int)

			for name, character := range characterMap {
				characterIdMap[name] = character.Id
			}

			return characterIdMap
		}

		characterListService := characters.NewCharacterListService(characterServiceMock, characterMapRepoMock)

		Convey("GetCharacterList Function", func() {

			Convey("When valid state", func() {

				characterList := characterListService.GetCharacterList()

				Convey("Should match map length", func() {
					So(len(characterList.Characters), ShouldEqual, 2)
				})
			})

			Convey("When invalid state", func() {

				Convey("When nil characterMap", func() {

					characterMapRepoMock.GetCharacterMapMock = func() map[string]int {
						return nil
					}

					characterList := characterListService.GetCharacterList()

					Convey("Should return empty character list", func() {
						So(len(characterList.Characters), ShouldEqual, 0)
					})

				})
			})

		})

	})

}
