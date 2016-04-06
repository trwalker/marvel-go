package rest

import (
	"net/http"
	"time"
)

type RestClientAdapter interface {
	Get(url string, timeout time.Duration) (resp *http.Response, body string, err error)
}
