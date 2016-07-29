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
	fmt.Println("URL:", "http://0.0.0.0:9000/")

	http.ListenAndServe("0.0.0.0:9000", apiHandler)
}
