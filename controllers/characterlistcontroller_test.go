package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type CharListCtrlTestContext struct {
	Controller CharacterListController
	Req        *http.Request
	Res        *httptest.ResponseRecorder
}

var charListControllerTestContext CharListCtrlTestContext = CharListCtrlTestContext{}

func (context *CharListCtrlTestContext) Setup() {
	context.Controller = CharacterListController{}
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
