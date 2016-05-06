package rest

import (
	"io/ioutil"
	"net/http"
	"time"
)

var RestClientAdapterInstance RestClientAdapter = &restClientAdapterImpl{}

type restClientAdapterImpl struct {
}

func NewRestClientAdapter() RestClientAdapter {
	restClientAdapter := &restClientAdapterImpl{}

	return restClientAdapter
}

func (restClientAdapter *restClientAdapterImpl) Get(url string, timeout time.Duration) (resp *http.Response, body string, err error) {
	resp = nil
	body = ""
	err = nil

	client := createClient(timeout)

	var req *http.Request
	req, err = createRequest(url)

	if err != nil {
		return
	}

	resp, err = sendRequest(client, req)

	if err != nil {
		return
	}

	body, err = readBody(resp)

	return
}

func createClient(timeout time.Duration) *http.Client {
	client := &http.Client{
		Timeout: timeout,
	}

	return client
}

func createRequest(url string) (req *http.Request, err error) {
	req, err = http.NewRequest("GET", url, nil)

	if err != nil {
		return
	}

	return
}

func sendRequest(client *http.Client, req *http.Request) (resp *http.Response, err error) {
	resp, err = client.Do(req)

	return
}

func readBody(resp *http.Response) (string, error) {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return string(body), err
}
