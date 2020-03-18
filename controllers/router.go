package controllers

import (
	"net/http"
	"regexp"

	"github.com/rjNemo/go-wiki/data"
)

// Router dispatch the request to the corresponding route handlers.
func Router(ctx data.Context) {
	hh := HomeHandler{Ctx: ctx}
	ph := PageHandler{Ctx: ctx}
	// uh := UserHandler{Users: UserStore}

	http.HandleFunc("/view/", makeHandler(ph.view))
	http.HandleFunc("/edit/", makeHandler(ph.edit))
	http.HandleFunc("/save/", makeHandler(ph.save))
	http.HandleFunc("/new/", ph.new)
	http.HandleFunc("/contact/", hh.contact)
	http.HandleFunc("/", hh.home)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
