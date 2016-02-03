package controllers

import (
	"github.com/trwalker/marvel-go/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CharListCtrlTestContext struct {
	Controller CharacterListController
	Req        *http.Request
	Res        *httptest.ResponseRecorder
}

type CharListServiceMock struct {
}

func (characterListService *CharListServiceMock) GetCharacterList() models.CharacterListModel {
	return models.CharacterListModel{}
}

var charListControllerTestContext *CharListCtrlTestContext = new(CharListCtrlTestContext)

func (context *CharListCtrlTestContext) Setup() {
	context.Controller = CharacterListController{
		CharacterListServiceInterface: &CharListServiceMock{},
	}

	context.Req = &http.Request{}
	context.Res = httptest.NewRecorder()
}

func (context *CharListCtrlTestContext) TearDown() {
	context = nil
}

func TestCharacterListControllerSuccess(t *testing.T) {
	charListControllerTestContext.Setup()
	defer charListControllerTestContext.TearDown()

	charListControllerTestContext.Controller.Get(charListControllerTestContext.Res, charListControllerTestContext.Req)
}
