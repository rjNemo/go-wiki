package main

import (
	"fmt"
	"time"

	"github.com/rjNemo/go-wiki/controller"
)

func main() {
	fmt.Printf("Start Go-wiki server on http://localhost:%s at %s\n", controller.Port, time.Now())
	controller.RegisteredRoutes()
}
