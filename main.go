package main

import (
	"log"
	"net/http"

	"github.com/rjNemo/go-wiki/data"
)

func main() {
	data.UsePSQL()
	// startServer(settings.Port, controller.Router)
}

func startServer(p string, r func()) {
	log.Printf("Start Go-wiki server on http://localhost:%s", p)
	port := ":" + p
	r()
	log.Fatal(http.ListenAndServe(port, nil))
}

// appBuilder
