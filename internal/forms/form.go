package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Form creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// Valid return true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has check if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}

// Required check any input field is empty or not
func (f *Form) Required(field ...string) {
	for _, input := range field {
		value := f.Get(input)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(input, "This field cannot be empty!!!")
		}
	}
}

// MinLength checks for string minium length
func (f *Form) MinLength(field string, leng int, r *http.Request) bool {
	value := f.Get(field)
	if len(value) < leng {
		f.Errors.Add(field, fmt.Sprintf("This field musy be at least %d chars", leng))
		return false
	}
	return true
}
