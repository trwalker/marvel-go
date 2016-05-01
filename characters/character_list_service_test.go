package characters

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"sync"
)

func TestCharacterListServiceSpec(t *testing.T) {

	Convey("CharacterListService Tests", t, func() {

		CharacterServiceMockInstance.ResetMock()
		CharacterMapRepoMockInstance.ResetMock()

		var characterListService CharacterListService = &CharacterListServiceImpl{
			CharacterServiceInterface: CharacterServiceMockInstance,
			CharacterMapRepoInterface: CharacterMapRepoMockInstance,
			lock: &sync.Mutex{},
			characterList: &CharacterListModel{
				Characters: make([]*CharacterModel, 0),
			},
		}

		Convey("GetCharacterList Function", func() {

			Convey("When valid state", func() {

				characterList := characterListService.GetCharacterList()

				Convey("Should match map length", func() {
					So(len(characterList.Characters), ShouldEqual, len(CharacterMapRepoMockInstance.characterMap))
				})

				Convey("Should contain map characters", func() {
					for _, listChar := range characterList.Characters {
						charId := CharacterMapRepoMockInstance.characterMap[listChar.Name]
						
						So(listChar.Id, ShouldEqual, charId)
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
