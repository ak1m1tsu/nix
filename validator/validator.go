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
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
)

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
