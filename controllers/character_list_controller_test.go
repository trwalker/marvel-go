package controllers

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/trwalker/marvel-go/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

var characterListController CharacterListController

var characterList *models.CharacterListModel

type CharacterListServiceMock struct {
}

func (characterListService *CharacterListServiceMock) GetCharacterList() *models.CharacterListModel {
	return characterList
}

func TestCharacterListControllerSpec(t *testing.T) {
	Convey("CharacterListController Tests", t, func() {

		CharacterListControllerInstance = initCharacterListController()

		characterList = &models.CharacterListModel{}
		characterList.Characters = append(characterList.Characters, &models.CharacterModel{
			Name:  "spider-man",
			Id:    1234,
			Image: "http://i.annihil.us/u/prod/marvel/foo.jpg",
		})

		characterListController = CharacterListController{
			CharacterListServiceInterface: &CharacterListServiceMock{},
		}

		req := &http.Request{}
		res := httptest.NewRecorder()

		Convey("Get Function", func() {

			Convey("When valid state", func() {

				characterListController.Get(res, req)
				model := &models.CharacterListModel{}
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

					characterList = &models.CharacterListModel{}
					characterListController.Get(res, req)
					model := &models.CharacterListModel{}
					json.Unmarshal([]byte(res.Body.String()), model)

					Convey("Should return empty character list", func() {
						So(len(model.Characters), ShouldEqual, 0)
					})

				})

			})

		})

	})

}
