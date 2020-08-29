package validator

import "github.com/go-playground/validator/v10"

// Validator has a group of methods to validate data.
type Validator interface {
	// Struct is used to validate a struct.
	Struct(s interface{}) error
}

// NewValidator initializes a new Validator using validator/v10.
func NewValidator() Validator {
	return validator.New()
}
