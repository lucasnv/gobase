package errors

// General erros goes from 0 to 99
const (
	UNKNOWN_ERROR_CODE            = 0
	INVALID_VALUE_ERROR_CODE      = 1
	NOT_FOUND_ERROR_CODE          = 2
	UNEXPECTED_COMMAND_ERROR_CODE = 3
)

type ErrorList []map[string]string

type ErrorApp struct {
	Code    uint16
	Message string
	Fields  ErrorList
}

func New(code uint16, msg string, fields ErrorList) error {
	return &ErrorApp{
		Code:    code,
		Message: msg,
		Fields:  fields,
	}
}

func NewUnknown() error {
	return New(UNKNOWN_ERROR_CODE, "unknown error", ErrorList{})
}

func NewInvalidValue(errorList ErrorList) error {
	return New(INVALID_VALUE_ERROR_CODE, "invalid value error", errorList)
}

func NewNotFound() error {
	return New(NOT_FOUND_ERROR_CODE, "not found error", ErrorList{})
}

func NewUnexpectedCommand() error {
	return New(UNEXPECTED_COMMAND_ERROR_CODE, "unexpected command", ErrorList{})
}

func (e *ErrorApp) Error() string {
	return e.Message
}

/*

func Unknown(err error) *errorhandler.ErrorHandler {
	return errorhandler.NewErrorHandler(UNKNOWN_ERROR_CODE, err)
}

func InvalidValue(field string, value string, err error) *errorhandler.ErrorHandler {
	return errorhandler.NewErrorHandler(INVALID_VALUE_ERROR_CODE, err)
}

func NotFound() *errorhandler.ErrorHandler {
	return errorhandler.NewErrorHandler(NOT_FOUND_ERROR_CODE, errors.New("not found"))
}

func UnexpectedCommand() *errorhandler.ErrorHandler {
	return errorhandler.NewErrorHandler(UNEXPECTED_COMMAND_ERROR_CODE, errors.New("unexpected command"))
}
*/
