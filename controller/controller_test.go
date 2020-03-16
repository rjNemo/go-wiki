package controller

import (
	"io"
	"net/http"
	"testing"

	"github.com/rjNemo/go-wiki/service"
)

func TestParseContactForm(t *testing.T) {
	var i io.Reader
	r, err := http.NewRequest(http.MethodPost, "/contact/post/", i)
	if err != nil {
		t.Errorf("%s", err)
	}
	ans := parseContactForm(r)
	if ans != service.NewMail("", "") {
		t.Errorf("parseContactForm(r) = %v", ans)
	}
}
