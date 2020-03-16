package controllers

import (
	"net/http"

	"github.com/rjNemo/go-wiki/models"
	"github.com/rjNemo/go-wiki/views"
)

// type PageHandler struct {
// 	Pages data.PageStore
// }

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := models.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	views.Template(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := models.LoadPage(title)
	if err != nil {
		p = models.NewPage(title, nil)
	}
	views.Template(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := models.NewPage(title, []byte(body))
	err := p.Save()
	checkError(err, w)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
