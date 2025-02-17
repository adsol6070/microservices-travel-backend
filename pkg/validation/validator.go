package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Validator instance
var validate = validator.New()

// ValidateStruct validates any struct using the validator package.
func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			var errorMessages []string
			for _, fieldError := range validationErrors {
				errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' failed validation: %s", fieldError.Field(), fieldError.Tag()))
			}
			return errors.New(fmt.Sprintf("Validation errors: %v", errorMessages))
		}
		return err
	}
	return nil
}
