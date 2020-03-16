package main

import (
	"log"
	"net/http"

	"github.com/rjNemo/go-wiki/controllers"
	"github.com/rjNemo/go-wiki/data"
	"github.com/rjNemo/go-wiki/settings"
)

func main() {
	data.UsePSQL()
	startServer(settings.Port, controllers.Router)
}

func startServer(p string, r func()) {
	log.Printf("Start Go-wiki server on http://localhost:%s", p)
	port := ":" + p
	r()
	log.Fatal(http.ListenAndServe(port, nil))
}

// appBuilder
