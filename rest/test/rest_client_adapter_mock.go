package resttest

import (
	"net/http"
	"time"
)

type RestClientAdapterMock struct {
	GetMock func(url string, timeout time.Duration) (resp *http.Response, body string, err error)
}

func (restClientAdapterMock *RestClientAdapterMock) Get(url string, timeout time.Duration) (resp *http.Response, body string, err error) {
	return restClientAdapterMock.GetMock(url, timeout)
}
