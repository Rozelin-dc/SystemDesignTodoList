package util

import (
	"regexp"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

var expressions = map[string]*regexp.Regexp{
	"userName":  regexp.MustCompile(`^[a-zA-Z0-9_-]{1,30}$`),
	"taskName":  regexp.MustCompile(`^.{1,60}$`),
	"timeLimit": regexp.MustCompile(`^[1-2][0-9]{3}/[0-1][0-9]/[0-3][0-9]$`),
	"password":  regexp.MustCompile(`^[\x21-\x7E]{10,30}$`),
}

func (v *CustomValidator) Validate(i interface{}) error {
	for key, e := range expressions {
		re := e
		v.validator.RegisterValidation(key, func(fl validator.FieldLevel) bool {
			return re.MatchString(fl.Field().String())
		})
	}
	v.validator.RegisterValidation("status", func(fl validator.FieldLevel) bool {
		return fl.Field().Int() == 0 || fl.Field().Int() == 1
	})

	return v.validator.Struct(i)
}

func GetValidator() echo.Validator {
	v := validator.New()

	return &CustomValidator{validator: v}
}
