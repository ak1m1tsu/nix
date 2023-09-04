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
//	    "your-module-path/validator"
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
//
//	    if err := validator.Validate(user); err != nil {
//	        fmt.Println("Validation failed:", err)
//	    } else {
//	        fmt.Println("Validation successful!")
//	    }
//	}
package validator

import (
	"errors"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ent "github.com/go-playground/validator/v10/translations/en"
)

var (
	// validate is an instance of the validator.Validate used for validation checks.
	validate *validator.Validate

	// uni is an instance of the ut.UniversalTranslator used for translation of error messages.
	uni *ut.UniversalTranslator

	// trans is an instance of the ut.Translator used for translation of error messages to English.
	trans ut.Translator
)

// init initializes the package by setting up the validator and translator instances.
// It registers the default translations for English language.
func init() {
	if validate == nil {
		validate = validator.New()
	}
	if uni == nil {
		en := en.New()
		uni = ut.New(en, en)
	}
	if trans == nil {
		trans, _ = uni.GetTranslator("en")
		_ = ent.RegisterDefaultTranslations(validate, trans)
	}
}

// Validate validates the input 'v' using the validator instance and
// returns an error if the validation fails. The error message includes all the
// validation errors separated by commas.
func Validate(v any) error {
	err := validate.Struct(v)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		errMsgs := make([]string, len(errs))
		for i, e := range errs {
			errMsgs[i] = e.Translate(trans)
		}
		return errors.New(strings.Join(errMsgs, ", "))
	}
	return nil
}

// RegisterRules registers the custom validation rules for the given 'v'
func RegisterRules(v any, rules map[string]string) {
	validate.RegisterStructValidationMapRules(rules, v)
}
