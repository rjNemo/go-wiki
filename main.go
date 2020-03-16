package main

import (
	"log"
	"net/http"

	"github.com/rjNemo/go-wiki/controller"
	"github.com/rjNemo/go-wiki/settings"
)

func main() {
	startServer(settings.Port, controller.Router)
}

func startServer(p string, r func()) {
	log.Printf("Start Go-wiki server on http://localhost:%s", p)
	port := ":" + p
	r()
	log.Fatal(http.ListenAndServe(port, nil))
}
