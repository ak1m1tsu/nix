// Package validator provides functions for validating data using the Go Playground Validator.
// It allows you to validate structs and individual fields, providing error messages in case of validation failures.
//
// The package relies on the go-playground/validator/v10 library for validation and go-playground/universal-translator
// for translating error messages to human readable text.
//
// Example Usage:
//
//	package main
//
//	import (
//	    "fmt"
//	    "github.com/romankravchuk/nix/validator"
//	)
//
//	type User struct {
//	    Username string `validate:"required,min=4,max=32"`
//	    Email    string `validate:"required,email"`
//	    Age      int    `validate:"gte=18"`
//	}
//
//	func main() {
//	    user := User{
//	        Username: "john_doe",
//	        Email:    "john.doe@example.com",
//	        Age:      25,
//	    }
//		v := validator.New()
//
//	    if isValid := v.Validate(user); !isValid {
//	        fmt.Println("Validation failed:", v.Errors())
//	    } else {
//	        fmt.Println("Validation successful!")
//	    }
//	}
package validator

import (
	"errors"
	"sync"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ent "github.com/go-playground/validator/v10/translations/en"
)

type Validator struct {
	validator  *validator.Validate
	translator ut.Translator

	errs   map[string]string
	errsMu sync.RWMutex
}

func New() *Validator {
	var (
		v        = validator.New()
		entrans  = en.New()
		trans, _ = ut.New(entrans, entrans).GetTranslator("en")
	)

	_ = ent.RegisterDefaultTranslations(v, trans)

	return &Validator{
		validator:  v,
		translator: trans,
		errs:       make(map[string]string),
	}
}

// Validate validates the given 'data' using the validator.
func (v *Validator) Validate(data interface{}) bool {
	var (
		errs validator.ValidationErrors
		err  = v.validator.Struct(data)
	)

	if err != nil {
		if errors.As(err, &errs) {
			v.errsMu.Lock()
			defer v.errsMu.Unlock()

			for _, e := range errs {
				v.errs[e.Field()] = e.Translate(v.translator)
			}
		}

		return false
	}

	return true
}

// Errors returns the validation errors as a map of field names to error messages.
func (v *Validator) Errors() map[string]string {
	v.errsMu.RLock()
	defer v.errsMu.RUnlock()

	return v.errs
}

// RegisterRules registers the custom validation rules for the given 'data'
func (v *Validator) RegisterRules(data any, rules map[string]string) {
	v.validator.RegisterStructValidationMapRules(rules, data)
}
