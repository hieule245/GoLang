package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
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
func (f *Form) Has(field string) bool {
	x := f.Get(field)
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
func (f *Form) MinLength(field string, leng int) bool {
	value := f.Get(field)
	if len(value) < leng {
		f.Errors.Add(field, fmt.Sprintf("This field musy be at least %d chars", leng))
		return false
	}
	return true
}

// IsEmail checks valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}

func (f *Form) IsPhone(field string) bool {
	pattern := `^(0[1-9][0-9]{8})$`
	match, err := regexp.MatchString(pattern, f.Get(field))
	if err != nil {
		f.Errors.Add(field, "Pattern is not valid!")
		return false
	}
	if !match {
		f.Errors.Add(field, "Invalid phone")
		return match
	}
	return match
}
