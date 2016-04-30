package characters

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var characterListService CharacterListService

func TestCharacterListServiceSpec(t *testing.T) {

	Convey("CharacterListService Tests", t, func() {

		characterMock := &CharacterModel{
			Id:          1,
			Name:        "spider-man",
			Description: "amazing spider man",
			Image:       "https://cdn.com/spidey.jpg",
		}

		CharacterServiceMockInstance.GetCharacterMockSetup(characterMock, true, nil)

		characterMapMock := make(map[string]int)
		characterMapMock["spider-man"] = 1
		characterMapMock["hulk"] = 2

		CharacterMapRepoMockInstance.GetCharacterMapMockSetup(characterMapMock)

		characterListService = &CharacterListServiceImpl{
			CharacterServiceInterface: CharacterServiceMockInstance,
			CharacterMapRepoInterface: CharacterMapRepoMockInstance,
			characterList: &CharacterListModel{
				Characters: make([]*CharacterModel, 0),
			},
		}

		Convey("GetCharacterList Function", func() {

			Convey("When valid state", func() {

				characterList := characterListService.GetCharacterList()

				Convey("Should match map length", func() {
					So(len(characterList.Characters), ShouldEqual, len(characterMapMock))
				})

				Convey("Should contain map characters", func() {
					for _, listChar := range characterList.Characters {
						_, found := characterMapMock[listChar.Name]
						So(found, ShouldBeTrue)
					}

				})
			})

			Convey("When invalid state", func() {

				Convey("When nil characterMap", func() {

					CharacterMapRepoMockInstance.GetCharacterMapMockSetup(nil)
					characterList := characterListService.GetCharacterList()

					Convey("Should return empty character list", func() {
						So(len(characterList.Characters), ShouldEqual, 0)
					})

				})
			})

		})

	})

}
