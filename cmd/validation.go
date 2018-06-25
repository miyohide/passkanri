package cmd

import (
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

func validateParams(p interface{}) error {
	errs := validate.Struct(p)
	return extractValidationErrors(errs)
}

func extractValidationErrors(err error) error {
	if err != nil {
		var errorText []string
		for _, err := range err.(validator.ValidationErrors) {
			errorText = append(errorText, validationErrorToText(err))
		}
		return fmt.Errorf("Parameter error: %s", strings.Join(errorText, "\n"))
	}
	return nil
}

func validationErrorToText(e validator.FieldError) string {
	f := e.Field()
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", f)
	}
	return fmt.Sprintf("%s is not valid %s", e.Field(), e.Value())
}
