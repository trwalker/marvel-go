package middleware

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ResponseHeadersTestContext struct {
	Req  *http.Request
	Res  *httptest.ResponseRecorder
	Body string
}

var responseHeadersTestContext *ResponseHeadersTestContext = new(ResponseHeadersTestContext)

func (context *ResponseHeadersTestContext) Setup() {
	context.Req = &http.Request{
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
	}

	context.Res = httptest.NewRecorder()
	context.Body = "{ \"foo\":\"bar\"}"
}

func (context *ResponseHeadersTestContext) TearDown() {
	context = nil
}

func (context *ResponseHeadersTestContext) SendApiRequest() {
	ResponseHeaders(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, context.Body)
	})).ServeHTTP(context.Res, context.Req)
}

func TestResponseHeadersHandlerContentType(t *testing.T) {
	responseHeadersTestContext.Setup()
	defer responseHeadersTestContext.TearDown()

	responseHeadersTestContext.SendApiRequest()

	contentType := responseHeadersTestContext.Res.HeaderMap.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Wrong \"Content-Type\" header: %s", contentType)
	}
}
