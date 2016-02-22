package controllers

import (
	"net/http"
)

type CharacterController interface {
	Get(res http.ResponseWriter, req *http.Request)
}
