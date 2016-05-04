package charservices

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/trwalker/marvel-go/characters/models"
	"sync"
	"testing"
)

func TestCharacterListServiceSpec(t *testing.T) {

	Convey("CharacterListService Tests", t, func() {

		characters := make(map[string]*charmodels.CharacterModel)
		characters["spider-man"] = &charmodels.CharacterModel{
			Id:          1,
			Name:        "spider-man",
			Description: "amazing spider man",
			Image:       "https://cdn.com/spidey.jpg",
		}

		characters["hulk"] = &charmodels.CharacterModel{
			Id:          2,
			Name:        "hulk",
			Description: "don't make me angry",
			Image:       "https://cdn.com/hulk.jpg",
		}

		var getCharacterError error = nil

		CharacterServiceMockInstance.GetCharacterMock = func(name string) (character *charmodels.CharacterModel, found bool, err error) {
			character, found = characters[name]
			err = getCharacterError

			return
		}

		CharacterMapRepoMockInstance.GetCharacterMapMock = func() map[string]int {
			characterMap := make(map[string]int)

			for name, character := range characters {
				characterMap[name] = character.Id
			}

			return characterMap
		}

		var characterListService CharacterListService = &CharacterListServiceImpl{
			CharacterServiceInterface: CharacterServiceMockInstance,
			CharacterMapRepoInterface: CharacterMapRepoMockInstance,
			lock: &sync.Mutex{},
			characterList: &charmodels.CharacterListModel{
				Characters: make([]*charmodels.CharacterModel, 0),
			},
		}

		Convey("GetCharacterList Function", func() {

			Convey("When valid state", func() {

				characterList := characterListService.GetCharacterList()

				Convey("Should match map length", func() {
					So(len(characterList.Characters), ShouldEqual, 2)
				})
			})

			Convey("When invalid state", func() {

				Convey("When nil characterMap", func() {

					CharacterMapRepoMockInstance.GetCharacterMapMock = func() map[string]int {
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
