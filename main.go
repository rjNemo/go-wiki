package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rjNemo/go-wiki/controllers"
	"github.com/rjNemo/go-wiki/data"
	"github.com/rjNemo/go-wiki/settings"
)

func main() {
	log.Println("*** Go-wiki v0.1 ©2020 ***")

	// connect to db
	// db, err := data.NewDB(settings.ConnStr)
	db, err := data.NewDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	// register store and inject them in handlers
	ctx := data.NewContext(db)

	// Migrate db …
	ctx.Pages.CreateTable()
	ctx.Users.CreateTable()

	// create handlers around context
	controllers.Router(ctx)

	// startServer
	log.Printf("Start Go-wiki server on http://localhost:%s", settings.Port)
	port := ":" + settings.Port
	log.Fatal(http.ListenAndServe(port, nil))
}
