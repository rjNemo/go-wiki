package controllers

import (
	"net/http"

	"github.com/rjNemo/go-wiki/data"
	"github.com/rjNemo/go-wiki/models"
	"github.com/rjNemo/go-wiki/views"
)

// PageHandler will respond to requests using Handlers
type PageHandler struct {
	Pages data.PageStore
} //and add each actual handler as a method

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := models.LoadPage(title)
	// p, err := ph.Pages.Get(1)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	views.Template(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := models.LoadPage(title)
	if err != nil {
		p = models.NewPage(0, title, nil)
	}
	views.Template(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := models.NewPage(0, title, []byte(body))
	err := p.Save()
	checkError(err, w)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
