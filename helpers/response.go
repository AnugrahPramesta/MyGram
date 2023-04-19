package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ResponseDefault struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

// APIResponse is for generating template responses
func APIResponse(message string, code int, err string) ResponseDefault {

	return ResponseDefault{
		Code:    code,
		Message: message,
		Error:   err,
	}
}

// FormatValidationError func which holds errors during user input validation
func FormatValidationError(err error) string {
	var errors string

	for _, e := range err.(validator.ValidationErrors) {
		if errors != "" {
			errors = fmt.Sprintf("%v, ", strings.TrimSpace(errors))
		}

		if e.Tag() == "email" {
			errors = errors + fmt.Sprintf("%v must be type %v", e.Field(), e.Tag())
		} else {
			errors = errors + fmt.Sprintf("%v is %v %v", e.Field(), e.Tag(), e.Param())
		}

		if e.Param() != "" && e.Type().Name() == "string" {
			errors = errors + " character"
		}

	}

	return strings.TrimSpace(errors) + "."
}
