package handler

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func validateReq(s interface{}) error {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	return validate.Struct(s)
}
