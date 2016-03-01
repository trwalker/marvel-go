package controllers

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/trwalker/marvel-go/characters"
	"net/http"
	"net/http/httptest"
	"testing"
)

var characterListController CharacterListController

var characterList *characters.CharacterListModel

type CharacterListServiceMock struct {
}

func (characterListService *CharacterListServiceMock) GetCharacterList() *characters.CharacterListModel {
	return characterList
}

func TestCharacterListControllerSpec(t *testing.T) {
	Convey("CharacterListController Tests", t, func() {

		characterList = &characters.CharacterListModel{}
		characterList.Characters = append(characterList.Characters, &characters.CharacterModel{
			Name:  "spider-man",
			Id:    1234,
			Image: "http://i.annihil.us/u/prod/marvel/foo.jpg",
		})

		characterListController = &CharacterListControllerImpl{
			CharacterListServiceInterface: &CharacterListServiceMock{},
		}

		req := &http.Request{}
		res := httptest.NewRecorder()

		Convey("Get Function", func() {

			Convey("When valid state", func() {

				characterListController.Get(res, req)
				model := &characters.CharacterListModel{}
				json.Unmarshal([]byte(res.Body.String()), model)

				Convey("Should write character list JSON with 1 item", func() {
					So(len(model.Characters), ShouldEqual, 1)
				})

				Convey("Should write character list JSON with spiderman", func() {
					So(model.Characters[0].Name, ShouldEqual, "spider-man")
				})
			})

			Convey("When invalid state", func() {

				Convey("When characterListModel is nil", func() {

					characterList = nil
					characterListController.Get(res, req)

					Convey("Should return empty character list", func() {
						So(res.Code, ShouldEqual, 404)
					})

				})

				Convey("When characterListModel is empty", func() {

					characterList = &characters.CharacterListModel{}
					characterListController.Get(res, req)
					model := &characters.CharacterListModel{}
					json.Unmarshal([]byte(res.Body.String()), model)

					Convey("Should return empty character list", func() {
						So(len(model.Characters), ShouldEqual, 0)
					})

				})

			})

		})

	})

}
