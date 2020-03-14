package model

import (
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func BlankPage() *Page {
	return &Page{Title: "Empty page", Body: []byte("Write some content")}
}

func NewPage(title string, body []byte) *Page {
	return &Page{Title: title, Body: body}
}

func (p *Page) Save() error {
	fileName := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(fileName, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	fileName := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
