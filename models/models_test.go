package models

import "testing"

func TestBlankPage(t *testing.T) {
	ans := BlankPage()
	if ans.title != "Empty page" && string(ans.body) != "Write some content" {
		t.Errorf("BlankPage() = %v; want &Page{title: 'Empty page', body: []byte('Write some content')}", ans)
	}
}

func TestNewPage(t *testing.T) {
	ans := NewPage("Test Page", []byte("This is a sample page"))
	if ans.title != "Test Page" && string(ans.body) != "This is a sample page" {
		t.Errorf("NewPage() = %v; want &Page{title: 'Test Page', body: []byte('This is a sample page')}", ans)
	}
}
