package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/hoge", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/hoge", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "aa")
	postedData.Add("c", "aaa")

	r, _ = http.NewRequest("POST", "/hoge", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("form does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/hoge", nil)
	form := New(r.PostForm)

	has := form.Has("a")
	if has {
		t.Error("form shows has when it should not")
	}

	postedData := url.Values{}
	postedData.Add("a", "abcde")

	r, _ = http.NewRequest("POST", "/hoge", nil)
	form = New(postedData)
	has = form.Has("a")

	if !has {
		t.Error("form does not show has when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/hoge", nil)
	form := New(r.PostForm)

	form.MinLength("a", 3)
	if form.Valid() {
		t.Error("form shows isMin when it should not")
	}

	isError := form.Errors.Get("a")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedData := url.Values{}
	postedData.Add("a", "abcde")
	form = New(postedData)

	form = New(postedData)
	form.MinLength("a", 100)

	if form.Valid() {
		t.Error("form does not show isMin when it should")
	}

	form = New(postedData)
	form.MinLength("a", 1)

	if !form.Valid() {
		t.Error("form does not show isMin when it should")
	}

	isError = form.Errors.Get("a")
	if isError != "" {
		t.Error("should not have an error, but got one")
	}

}

func TestIsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/hoge", nil)
	form := New(r.PostForm)

	isEmail := form.IsEmail("a")
	if isEmail {
		t.Error("form shows email when it should not")
	}

	postedData := url.Values{}
	postedData.Add("a", "t@gmail.com")
	r.PostForm = postedData
	form = New(r.PostForm)
	isEmail = form.IsEmail("a")

	if !isEmail {
		t.Error("form does not show email when it should")
	}

}
