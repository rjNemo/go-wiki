package controllers

import (
	"log"
	"net/http"

	"github.com/rjNemo/go-wiki/services"
	"github.com/rjNemo/go-wiki/views"
)

// type HomeHandler struct {
// }

func homeHandler(w http.ResponseWriter, r *http.Request) {
	views.Template(w, "home", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
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
