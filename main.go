package main

import (
	"log"
	"net/http"

	"github.com/rjNemo/go-wiki/controllers"
	"github.com/rjNemo/go-wiki/data"
	"github.com/rjNemo/go-wiki/settings"
)

func main() {
	// connect to db
	db, err := data.NewDB(settings.ConnStr)
	if err != nil {
		log.Fatal(err)
	}

	// register store and inject them in handlers
	ctx := data.NewContext(db)
	ctx.Pages.CreateTable()
	// ctx.Users.CreateTable()

	ph := controllers.PageHandler{Ctx: ctx}
	// uh := controllers.UserHandler{Users: UserStore}

	// startServer(, controllers.Router)
	log.Printf("Start Go-wiki server on http://localhost:%s", settings.Port)
	port := ":" + settings.Port
	controllers.Router(ph)
	log.Fatal(http.ListenAndServe(port, nil))
}

// func startServer(p string, r func()) {
// }

// // appBuilder
