package validator

import (
	"reflect"
	"strings"

	"github.com/MrTj458/markednotes"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func New() Validator {
	v := validator.New()

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return Validator{v}
}

func (v Validator) Struct(s any) ([]markednotes.ErrorField, bool) {
	err := v.validate.Struct(s)
	if err != nil {
		errors := []markednotes.ErrorField{}
		for _, err := range err.(validator.ValidationErrors) {
			newErr := markednotes.ErrorField{
				Name:   err.Field(),
				Detail: "invalid " + err.Tag(),
			}

			if len(err.Param()) > 0 {
				newErr.Detail = err.Tag() + "=" + err.Param()
			}

			errors = append(errors, newErr)
		}
		return errors, false
	}
	return nil, true
}
