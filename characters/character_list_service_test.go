package characters

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var characterListService CharacterListService

var characterMapMock map[string]*CharacterModel

type CharacterMapRepoMock struct {
}

func (charMapRepoMock *CharacterMapRepoMock) GetCharacterMap() map[string]*CharacterModel {
	return characterMapMock
}

func TestCharacterListServiceSpec(t *testing.T) {

	Convey("CharacterListService Tests", t, func() {

		characterMapMock = make(map[string]*CharacterModel)

		characterMapMock["spider-man"] = &CharacterModel{
			Id:    1,
			Name:  "spider-man",
			Image: "http://i.annihil.us/u/prod/marvel/bar.jpg",
		}

		characterMapMock["hulk"] = &CharacterModel{
			Id:    2,
			Name:  "hulk",
			Image: "http://i.annihil.us/u/prod/marvel/foo.jpg",
		}

		characterListService = &CharacterListServiceImpl{
			CharacterMapRepoInterface: &CharacterMapRepoMock{},
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

					characterMapMock = nil
					characterList := characterListService.GetCharacterList()

					Convey("Should return empty character list", func() {
						So(len(characterList.Characters), ShouldEqual, 0)
					})

				})
			})

		})

	})

}
