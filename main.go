package main

import (
	"log"
	"net/http"

	"github.com/rjNemo/go-wiki/controller"
	"github.com/rjNemo/go-wiki/service"
)

func main() {
	service.PaymentIntent(1000, "ruidy.nemausat@gmail.com")
	startServer(controller.Port, controller.Router)
}

func startServer(p string, r func()) {
	log.Printf("Start Go-wiki server on http://localhost:%s\n", controller.Port)
	port := ":" + p
	r()
	log.Fatal(http.ListenAndServe(port, nil))
}
