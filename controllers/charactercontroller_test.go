package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type CharCtrlTestContext struct {
	Controller CharacterController
	Req        *http.Request
	Res        *httptest.ResponseRecorder
}

var charControllerTestContext CharCtrlTestContext = CharCtrlTestContext{}

func (context *CharCtrlTestContext) Setup() {
	context.Controller = CharacterController{}
	context.Req = &http.Request{}
	context.Res = httptest.NewRecorder()
}

func (context *CharCtrlTestContext) TearDown() {
	context = nil
}

func TestCharacterControllerSuccess(t *testing.T) {
	charControllerTestContext.Setup()
	charControllerTestContext.TearDown()

	charControllerTestContext.Controller.Get(charControllerTestContext.Res, charControllerTestContext.Req)
}
