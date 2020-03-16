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
	views.Template(w, "contact", nil)
}

func postContactHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	checkError(err, w) // bad error handling
	mail := parseContactForm(r)
	// mail.Send()
	log.Println(mail)
	views.Template(w, "contact_sent", nil)
}

func parseContactForm(r *http.Request) services.Mail {
	log.Println(r.PostForm)
	return services.NewMail(r.PostFormValue("email"), r.PostFormValue("message"))
}

func checkError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
