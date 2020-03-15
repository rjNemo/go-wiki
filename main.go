package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rjNemo/go-wiki/controller"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i, n := range numbers {
		fmt.Println("Slice item", i, "is", n)
	}

	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for k, val := range countryCapitalMap {
		fmt.Println("Capital of", k, "is", val)
	}
	// startServer(controller.Port)
}

func startServer(p string) {
	log.Printf("Start Go-wiki server on http://localhost:%s\n", controller.Port)
	port := ":" + p
	controller.RegisteredRoutes()
	log.Fatal(http.ListenAndServe(port, nil))
}
