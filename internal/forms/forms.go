package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	// Data holds the form data
	url.Values
	// Valid indicates if the form is valid or not
	Errors errors
}

// New initializes a new Form struct with the provided form data and an empty errors map
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Valid checks if the form is valid (i.e., no errors)
// Valid returns true if there are no errors in the form
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Has checks if a form field is present and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	if f.Get(field) == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}

// MinLength checks if a form field has a minimum length
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	value := f.Get(field)
	if len(value) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

// IsEmail checks if a form field contains a valid email address
func (f *Form) IsEmail(field string) bool {
	value := f.Get(field)
	if !govalidator.IsEmail(value) {
		f.Errors.Add(field, "Invalid email address")
		return false
	}
	return true
}
