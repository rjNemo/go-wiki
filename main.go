package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rjNemo/go-wiki/controller"
)

func main() {
	fmt.Printf("Start Go-wiki server on http://localhost:%s at %s\n", controller.Port, time.Now())
	startServer(controller.Port)
}

func startServer(p string) {
	port := ":" + p
	controller.RegisteredRoutes()
	log.Fatal(http.ListenAndServe(port, nil))
}
