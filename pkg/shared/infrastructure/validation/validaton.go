package validation

import (
	"errors"

	v10 "github.com/go-playground/validator/v10"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/validator"
)

type V10Validator struct {
	service *v10.Validate
}

type V10ValidationError struct {
	fieldValue string
	errorValue string
}

func (e V10ValidationError) Error() string {
	return "Validation error: " + e.errorValue
}

func (e V10ValidationError) Field() string {
	return e.fieldValue
}
//TODO: ver si tengo que devolver un puntero mejor
func New() validator.Validator {
	return V10Validator{
		service: v10.New(),
	}
}

func (v V10Validator) Var(field string, value interface{}, rules string) validator.ValidationError {
	err := v.service.Var(value, rules)
	var validatorErr v10.ValidationErrors

	if err != nil && errors.As(err, &validatorErr) {
		return V10ValidationError{fieldValue: field, errorValue: getErrorMessage(validatorErr)}
	}

	return nil
}

func getErrorMessage(err v10.ValidationErrors) string {
	var message string

	for _, e := range err {
		message = e.Tag()
	}

	return message
}
