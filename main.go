package main

import (
	"fmt"
	"github.com/trwalker/marvel-go/config/handlers"
	"net/http"
)

func main() {
	apiHandler := handlers.BuildApiHandler()

	startServer(apiHandler)
}

func startServer(apiHandler http.Handler) {
	fmt.Println("Starting Web Server...")
	fmt.Println("URL:", "http://127.0.0.1:9000/")

	http.ListenAndServe("127.0.0.1:9000", apiHandler)
}
