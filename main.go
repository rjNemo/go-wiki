package main

import (
	"log"
	"net/http"

	"github.com/rjNemo/go-wiki/controller"
)

func main() {
	startServer(controller.Port)
}

func startServer(p string) {
	log.Printf("Start Go-wiki server on http://localhost:%s\n", controller.Port)
	port := ":" + p
	controller.RegisteredRoutes()
	log.Fatal(http.ListenAndServe(port, nil))
}
