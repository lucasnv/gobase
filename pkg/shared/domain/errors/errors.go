package errors

const (
	UNKNOWN_CODE             = 0
	BAD_REQUEST_CODE         = 10
	UNAUTHORIZED_CODE        = 11
	FORBIDDEN_CODE           = 12
	NOT_FOUND_CODE           = 13
	INNER_ERROR_CODE         = 14
	UNEXPECTED_COMMAND_ERROR = 15
	INVALID_UUID_ERROR       = 16
)

type ErrCode uint16

type ErrorsList map[string]string

// TODO: We have to take into account, this struct don't allow an errorList with more than one element
type AppError struct {
	ErrCode    ErrCode
	ErrMessage string
	ErrorsList ErrorsList
}

func NewAppError(c ErrCode) *AppError {
	return &AppError{
		ErrCode:    c,
		ErrMessage: getErrorMessage(c),
	}
}

func (e AppError) Code() ErrCode {
	return e.ErrCode
}

func (e AppError) Message() string {
	return e.ErrMessage
}

func (e AppError) Errors() ErrorsList {
	return e.ErrorsList
}

func (e AppError) Error() string {
	return e.ErrMessage
}

func (e *AppError) AddError(field string, message string) {
	e.ErrorsList = make(ErrorsList)
	e.ErrorsList[field] = message
}

func getErrorMessage(c ErrCode) string {
	switch c {
	case BAD_REQUEST_CODE:
		return "Server cannot process the request due to something wrong with the request."
	case FORBIDDEN_CODE:
		return "Client does not have access to the requested resource."
	case INNER_ERROR_CODE:
		return "Server encountered an unexpected condition."
	case NOT_FOUND_CODE:
		return "Server cannot find the requested resource."
	case UNAUTHORIZED_CODE:
		return "Client needs authentication to get requested resource"
	case UNEXPECTED_COMMAND_ERROR:
		return "Unexpected command."
	case INVALID_UUID_ERROR:
		return "Invalid Id."
	case UNKNOWN_CODE:
		return "Unknown error."
	default:
		return "Unknown error."
	}
}

type App interface {
	Code() ErrCode
	Message() string
	Errors() ErrorsList
	Error() string
}
