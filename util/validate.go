package util

import (
	"regexp"

	"gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	validator *validator.Validate
}

var expressions = map[string]*regexp.Regexp{
	"userName": regexp.MustCompile(`^[a-zA-Z0-9_-]{1,30}$`),
	"taskName": regexp.MustCompile(`^.{1,60}$`),
	"password": regexp.MustCompile(`^[\x20-\x7E]{10,32}$`),
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func GetValidator() *Validator {
	v := validator.New()

	for key, e := range expressions {
		v.RegisterValidation(key, func(fl validator.FieldLevel) bool {
			return e.MatchString(fl.Field().String())
		})
	}

	return &Validator{v}
}
