package helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	XValidator struct {
		Validator *validator.Validate
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func (v XValidator) validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	fmt.Println("checking if there is not errors", validationErrors)
	return validationErrors
}

func ValidateData(merch interface{}) []string {
	var myValidator = XValidator{Validator: validate}
	errs := myValidator.validate(merch)

	if len(errs) == 0 {
		return nil
	}

	var errMsgs []string
	for _, err := range errs {
		errMsgs = append(errMsgs, fmt.Sprintf(
			"field:'%s' with value:'%s' is '%s'",
			err.FailedField,
			err.Value,
			err.Tag,
		))
	}

	return errMsgs
}
