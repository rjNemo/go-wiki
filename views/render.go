package views

import (
	"log"
	"net/http"
	"text/template"

	"github.com/rjNemo/go-wiki/models"
	"github.com/rjNemo/go-wiki/settings"
)

func Template(w http.ResponseWriter, tmpl string, p *models.Page) {
	// err := templates.ExecuteTemplate(w, "templates/"+tmpl+".html", p)
	t, err := template.ParseFiles(getTmplName("base"), getTmplName(tmpl))

	if err != nil {
		log.Println(err.Error())
		return
	}

	err = t.Execute(w, p)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func getTmplName(tmpl string) string {
	return settings.TmplDir + tmpl + ".html"
}

// var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html")) // add slice of fileNAmes

// func ParseTemplates() *template.Template {
// 	return template.Must(template.ParseFiles("templates/edit.html", "templates/view.html")) // add slice of fileNAmes
// }
