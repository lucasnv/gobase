package response

type AppError struct {
	Code    int
	Message string
	Errors  interface{}
}

var (
	LOGIN_UNKNOWN = NewError(202, "LOGIN_UNKNOWN")
	LOGIN_ERROR   = NewError(203, "LOGIN_ERROR")
	VALID_ERROR   = NewError(300, "VALID_ERROR")
	BAD_REQUEST   = NewError(400, "BAD_REQUEST")
	UNAUTHORIZED  = NewError(401, "UNAUTHORIZED")
	NOT_FOUND     = NewError(404, "NOT_FOUND")
	INNER_ERROR   = NewError(500, "INNER_ERROR")
)

func (e *AppError) Error() string {
	return e.Message
}

func NewError(code int, msg string) *AppError {
	return &AppError{
		Message: msg,
		Code:    code,
	}
}
