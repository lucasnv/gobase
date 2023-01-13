package validation

import (
	"errors"

	"github.com/go-playground/validator/v10"
	internalError "[MODULE_URL]/pkg/shared/domain/errors"
)

type Validator interface {
	Var(field interface{}, tag string) error
	Struct(s interface{}) error
}

type Validation struct {
	service *validator.Validate
	field   string
}

func New() *Validation {
	service := validator.New()

	return &Validation{
		service: service,
	}
}

func (v *Validation) Var(field string, value interface{}, rules string) error {
	v.field = field
	err := v.service.Var(value, rules)
	var validatorErr validator.ValidationErrors

	if err != nil && errors.As(err, &validatorErr) {
		errorList := getFieldsWithErrors(validatorErr)

		return internalError.NewInvalidValue(errorList)
	}

	return nil
}

func (v *Validation) Struct(s interface{}) error {
	err := v.service.Struct(s)
	var validatorErr validator.ValidationErrors

	if err != nil && errors.As(err, &validatorErr) {
		errorList := getFieldsWithErrors(validatorErr)

		return internalError.NewInvalidValue(errorList)
	}

	return nil
}

func getFieldsWithErrors(err validator.ValidationErrors) internalError.ErrorList {
	InvalidFields := make(internalError.ErrorList, 0)

	for _, e := range err {
		errors := map[string]string{}
		errors[e.Field()] = e.Tag()
		InvalidFields = append(InvalidFields, errors)
	}

	return InvalidFields
}
