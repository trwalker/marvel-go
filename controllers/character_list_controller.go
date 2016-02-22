package controllers

import (
	"net/http"
)

type CharacterListController interface {
	Get(res http.ResponseWriter, req *http.Request)
}
