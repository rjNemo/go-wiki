package views

import (
	"net/http"
	"text/template"

	"github.com/rjNemo/go-wiki/model"
	"github.com/rjNemo/go-wiki/settings"
)

func Template(w http.ResponseWriter, tmpl string, p *model.Page) {
	// err := templates.ExecuteTemplate(w, "templates/"+tmpl+".html", p)
	t, err := template.ParseFiles(getTmplName("base"), getTmplName(tmpl))
	checkError(err, w)
	err = t.Execute(w, p)
	checkError(err, w)
}

func getTmplName(tmpl string) string {
	return settings.TmplDir + tmpl + ".html"
}

// var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html")) // add slice of fileNAmes

// func ParseTemplates() *template.Template {
// 	return template.Must(template.ParseFiles("templates/edit.html", "templates/view.html")) // add slice of fileNAmes
// }
