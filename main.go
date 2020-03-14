package main

import (
	"fmt"
	"time"

	"github.com/rjNemo/go-wiki/controller"
)

func main() {
	fmt.Printf("Start Go-wiki server on http://localhost:%d at %s\n", controller.Port, time.Now())
	controller.RegisteredRoutes()
}
