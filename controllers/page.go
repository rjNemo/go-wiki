package controllers

import (
	"log"
	"net/http"

	"github.com/rjNemo/go-wiki/data"
	"github.com/rjNemo/go-wiki/models"
	"github.com/rjNemo/go-wiki/views"
)

// PageHandler will respond to requests using Handlers
type PageHandler struct {
	Ctx data.Context
}

func (ph PageHandler) index(w http.ResponseWriter, r *http.Request) {
	index, err := ph.Ctx.Pages.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	views.Template(w, "index", struct{ Wikis []models.Page }{index})
}

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

	if !ph.Ctx.Pages.Exists(title) {
		p := models.NewPage(0, title, []byte(body))
		ph.Ctx.Pages.Add(*p)
	} else {
		p, err := ph.Ctx.Pages.Get(title)
		if err != nil {
			log.Fatal(err)
		}
		p.SetBody([]byte(body))
		ph.Ctx.Pages.Update(p.ID(), p)
	}

	// checkError(err, w)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func (ph PageHandler) new(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("newPage")
	http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	return
}
