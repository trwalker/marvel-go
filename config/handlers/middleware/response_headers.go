package middleware

import (
	"net/http"
)

func ResponseHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		handler.ServeHTTP(res, req)
	})
}
