package model

import (
	"io/ioutil"
)

// A Page own a wiki page data and has a title and a body.
type Page struct {
	title string
	body  []byte
}

// BlankPage constructor returns a pointer to a blank Page.
func BlankPage() *Page {
	return &Page{title: "Empty page", body: []byte("Write some content")}
}

// NewPage constructor returns a pointer to a sample Page.
func NewPage(title string, body []byte) *Page {
	return &Page{title: title, body: body}
}

// Save a page to the 'data/' folder in txt format.
func (p *Page) Save() error {
	fileName := "data/" + p.title + ".txt"
	return ioutil.WriteFile(fileName, p.body, 0600)
}

// LoadPage reads a saved page data and returns a pointer to the Page.
func LoadPage(title string) (*Page, error) {
	fileName := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{title: title, body: body}, nil
}
