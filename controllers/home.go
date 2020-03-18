package controllers

import (
	"log"
	"net/http"

	"github.com/rjNemo/go-wiki/data"
	"github.com/rjNemo/go-wiki/models"
	"github.com/rjNemo/go-wiki/services"
	"github.com/rjNemo/go-wiki/views"
)

// HomeHandler responds to requests using Handlers
type HomeHandler struct {
	Ctx data.Context
}

func (hh HomeHandler) home(w http.ResponseWriter, r *http.Request) {
	index, err := hh.Ctx.Pages.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	views.Template(w, "home", struct{ Wikis []models.Page }{index})
}

func (hh HomeHandler) contact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		views.Template(w, "contact", nil)
		return
	}
	mail := services.NewMail(r.PostFormValue("email"), r.PostFormValue("message"))
	log.Println(mail)
	views.Template(w, "contact", struct{ Success bool }{true})
}

func checkError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
