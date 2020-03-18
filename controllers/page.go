package controllers

import (
	"net/http"

	"github.com/rjNemo/go-wiki/data"
	"github.com/rjNemo/go-wiki/models"
	"github.com/rjNemo/go-wiki/views"
)

// PageHandler will respond to requests using Handlers
type PageHandler struct {
	Ctx data.Context
} //and add each actual handler as a method

// func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
func (ph PageHandler) view(w http.ResponseWriter, r *http.Request, title string) {
	// p, err := models.LoadPage(title)
	p, err := ph.Ctx.Pages.Get(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	views.Template(w, "view", p)
}

func (ph PageHandler) edit(w http.ResponseWriter, r *http.Request, title string) {
	// p, err := models.LoadPage(title)
	p, err := ph.Ctx.Pages.Get(title)
	if err != nil {
		p = *models.NewPage(0, title, nil)
	}
	views.Template(w, "edit", p)
}

func (ph PageHandler) save(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := models.NewPage(0, title, []byte(body))
	// err := p.Save()

	if !ph.Ctx.Pages.Exists(title) {
		ph.Ctx.Pages.Add(*p)
	} else {
		ph.Ctx.Pages.Update(p.ID(), *p)
	}

	// checkError(err, w)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
