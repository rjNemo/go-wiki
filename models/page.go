package models

import (
	"errors"
	"strings"
)

// A Page own a wiki page data and has a title and a body.
type Page struct {
	id    uint
	title string
	body  []byte
}

// ID exposes Page's title field
func (p Page) ID() uint {
	return p.id
}

// Title exposes Page's title field
func (p Page) Title() string {
	return p.title
}

// SetTitle edits Page's title field
func (p *Page) SetTitle(s string) {
	p.title = s
}

// Body exposes Page's body field
func (p Page) Body() []byte {
	return p.body
}

// SetBody edits Page's body field
func (p *Page) SetBody(b []byte) {
	p.body = b
}

// BlankPage constructor returns a pointer to a blank Page.
func BlankPage() *Page {
	return &Page{title: "Empty page", body: []byte("Write some content")}
}

// NewPage constructor returns a pointer to a sample Page.
// func NewPage(id uint, title string, body []byte) (*Page, error) {
func NewPage(id uint, title string, body []byte) *Page {
	// if id < 0 {
	// 	return nil, errNegID
	// }
	// if title == "" {
	// 	return nil, errBlankTitle
	// }
	return &Page{
		id:    id,
		title: strings.Title(title),
		body:  body}
	// , nil
}

var errBlankTitle = errors.New("The title must not be blank")
var errNegID = errors.New("The ID must be unsigned")

// Save a page to the 'data/' folder in txt format.
// func (p *Page) Save() error {
// 	fileName := "data/files/" + p.title + ".txt"
// 	return ioutil.WriteFile(fileName, p.body, 0600)
// }

// LoadPage reads a saved page data and returns a pointer to the Page.
// func LoadPage(title string) (*Page, error) {
// 	fileName := "data/files/" + title + ".txt"
// 	body, err := ioutil.ReadFile(fileName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Page{title: title, body: body}, nil
// }
