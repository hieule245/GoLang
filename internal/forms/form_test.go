package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("show does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("shows form does not have field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have an error but not have one")
	}

	postedData := url.Values{}
	postedData.Add("some_field", "some value")
	form = New(postedData)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows minlength of 100 met when data is shorter")
	}

	postedData = url.Values{}
	postedData.Add("another_field", "abc values")
	form = New(postedData)

	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("shows minlength of 1 is not met when data is")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should not have an error but got one")
	}
}

func TestForm_IsMail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	form.IsEmail("x")
	if form.Valid() {
		t.Error("form show valid email for non-existent field")
	}
	postedData := url.Values{}
	postedData.Add("email", "abc@gmail.com")
	form = New(postedData)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got invalid email form")
	}

	postedData = url.Values{}
	postedData.Add("another_email", "x")
	form = New(postedData)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got valid for the invalid email form")
	}
}

func TestForm_IsPhone(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	form.IsPhone("x")
	if form.Valid() {
		t.Error("form show valid email for non-existent field")
	}
	postedData = url.Values{}
	postedData.Add("phone", "0919280763")
	form = New(postedData)

	form.IsPhone("phone")
	if !form.Valid() {
		t.Error("got invalid phone form")
	}

	postedData = url.Values{}
	postedData.Add("phone", "091")
	form = New(postedData)

	form.IsPhone("phone")
	if form.Valid() {
		t.Error("valid form phone should have 10 numbers")
	}

	postedData = url.Values{}
	postedData.Add("phone", "1234567890")
	form = New(postedData)

	form.IsPhone("phone")
	if form.Valid() {
		t.Error("valid form phone should start with 0")
	}

	postedData = url.Values{}
	postedData.Add("phone", "0023456789")
	form = New(postedData)

	form.IsPhone("phone")
	if form.Valid() {
		t.Error("valid form phone shouldn't be 0 in the second number")
	}
}
