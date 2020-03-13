package main

import (
	"fmt"
	"time"

	"github.com/rjNemo/go-wiki/controllers"
)

func main() {
	fmt.Printf("Start Go-wiki server at %s\n", time.Now())
	controllers.RegisteredRoutes()
}
