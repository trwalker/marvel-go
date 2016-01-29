package services

import (
	"testing"
)

type CharListServiceTestContext struct {
	Service CharacterListService
}

var charListServiceTestContext *CharListServiceTestContext = new(CharListServiceTestContext)

func (context *CharListServiceTestContext) Setup() {
}

func (context *CharListServiceTestContext) TearDown() {

}

func TestCharacterListServiceValidState(t *testing.T) {

}
