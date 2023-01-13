package errorhandler

/*
type errorList []map[string]string

type ErrorHandler struct {
	Code    uint16
	Message string
	Fields  errorList
	Err     error
}

func NewErrorHandler(code uint16, err error) *ErrorHandler {

	var validatorErr validator.ValidationErrors
	fieldsWithErrors := make(errorList, 0)

	if errors.As(err, &validatorErr) {
		fieldsWithErrors = getFieldsWithErrors(validatorErr)
	}

	return &ErrorHandler{
		Code:    code,
		Message: err.Error(),
		Fields:  fieldsWithErrors,
		Err:     err,
	}
}

func (eh *ErrorHandler) Error() string {
	return eh.Message
}

func getFieldsWithErrors(err validator.ValidationErrors) errorList {
	InvalidFields := make(errorList, 0)

	for _, e := range err {
		errors := map[string]string{}
		errors[e.Field()] = e.Tag()
		InvalidFields = append(InvalidFields, errors)
	}

	return InvalidFields
}
*/
