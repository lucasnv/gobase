package errors

const (
	UNKNOWN_ERROR                 = 0
	BAD_REQUEST_ERROR             = 10
	UNAUTHORIZED_ERROR            = 11
	FORBIDDEN_ERROR               = 12
	NOT_FOUND_ERROR               = 13
	INNER_ERROR                   = 14
	UNEXPECTED_COMMAND_ERROR      = 15
	INVALID_UUID_ERROR            = 16
	MALFORMED_FILTER_ERROR        = 17
	INVALID_OPERATOR_FILTER_ERROR = 18
	INVALID_CRITERIA_ERROR        = 19
)

type ErrCode uint16

type ErrorsList map[string]string

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
	case BAD_REQUEST_ERROR:
		return "Server cannot process the request due to something wrong with the request."
	case FORBIDDEN_ERROR:
		return "Client does not have access to the requested resource."
	case INNER_ERROR:
		return "Server encountered an unexpected condition."
	case NOT_FOUND_ERROR:
		return "Server cannot find the requested resource."
	case UNAUTHORIZED_ERROR:
		return "Client needs authentication to get requested resource"
	case UNEXPECTED_COMMAND_ERROR:
		return "Unexpected command."
	case INVALID_UUID_ERROR:
		return "Invalid Id."
	case MALFORMED_FILTER_ERROR:
		return "Server cannot process the filter check the correct format. It must be something like: ?filter=[criteria]::[operator]::[parameters],[criteria]::[operator]::[parameters]."
	case INVALID_OPERATOR_FILTER_ERROR:
		return "Server cannot process the filter, you have use an invalid operator."
	case INVALID_CRITERIA_ERROR:
		return "Server cannot process the filter, you have use an invalid criteria."
	case UNKNOWN_ERROR:
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
