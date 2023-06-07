package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"strings"
)

// Form creates a custom form struct, embeds a url.Values object
type Form struct {
	Errors errors
}

// Valid return true if there are no errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initializes a form struct
func New() *Form {
	return &Form{
		errors(map[string][]string{}),
	}
}

// Required checks for required field
func (f *Form) Required(fields ...string) {
	for _, value := range fields {
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(value, "This field cannot be blank")
		}
	}
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string) bool {
	if field == "" {
		return false
	}
	return true
}

// MinLength checks for string minimum length
func (f *Form) MinLength(field string, length int) bool {
	if len(field) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

// IsEmail checks for valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(field) {
		f.Errors.Add(field, "Invalid email address")
	}
}
