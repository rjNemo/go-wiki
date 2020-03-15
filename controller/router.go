package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"

	"github.com/rjNemo/go-wiki/model"
	"github.com/rjNemo/go-wiki/service"
)

// func ParseTemplates() *template.Template {
// 	return template.Must(template.ParseFiles("templates/edit.html", "templates/view.html")) // add slice of fileNAmes
// }

// Router dispatch the request to the corresponding route handlers.
func Router() {
	// http.HandleFunc("/", loveHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/contact/", contactHandler)
	http.HandleFunc("/contact/post", postContactHandler)
	http.HandleFunc("/", homeHandler)
}

// func loveHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[1:]
// 	p := model.NewPage(title, nil) // already a pointer
// 	renderTemplate(w, "love", p)
// }

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home", nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := model.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := model.LoadPage(title)
	if err != nil {
		p = model.NewPage(title, nil)
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := model.NewPage(title, []byte(body))
	err := p.Save()
	checkError(err, w)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact", nil)
}

func postContactHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	checkError(err, w) // bad error handling
	mail := parseContactForm(r)
	service.MailClient(mail)
	fmt.Println(mail)
	renderTemplate(w, "contact_sent", nil)
}

func parseContactForm(r *http.Request) service.Mail {
	return service.NewMail(r.PostFormValue("email"), r.PostFormValue("message"))
}

// var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html")) // add slice of fileNAmes

func renderTemplate(w http.ResponseWriter, tmpl string, p *model.Page) {
	// err := templates.ExecuteTemplate(w, "templates/"+tmpl+".html", p)
	t, err := template.ParseFiles(getTmplName("base"), getTmplName(tmpl))
	checkError(err, w)
	err = t.Execute(w, p)
	checkError(err, w)
}

func getTmplName(tmpl string) string {
	return tmplDir + tmpl + ".html"
}

func checkError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

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
