package domain

import (
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/validator"
)

// Users' errors goes from 1000 to 1999
const (
	UNKNOWN_USER_ERROR    = 1000
	INVALID_USER_ERROR    = 1001
	REPOSITORY_USER_ERROR = 1002
	NOT_FOUND_ERROR       = 1003
	INVALID_USER_ID_ERROR = 1004
	USER_DELETE_ERROR     = 1005
)

func NewUserError(c errors.ErrCode, ve ...validator.ValidationError) *errors.AppError {
	var error errors.AppError = errors.AppError{
		ErrCode:    c,
		ErrMessage: getErrorMessage(c),
	}

	if ve != nil {
		for _, v := range ve {
			error.AddError(v.Field(), v.Error())
		}
	}

	return &error
}

func getErrorMessage(c errors.ErrCode) string {
	switch c {
	case INVALID_USER_ERROR:
		return "BC USER: cannot process the request due to something wrong with the data."
	case UNKNOWN_USER_ERROR:
		return "BC USER: Unknown user error."
	case REPOSITORY_USER_ERROR:
		return "BC USER: There was a problema in user repository."
	case NOT_FOUND_ERROR:
		return "BC USER: User not found."
	case INVALID_USER_ID_ERROR:
		return "BC USER: The user id has an invalid format."
	case USER_DELETE_ERROR:
		return "BC USER: The user could not be deleted."
	default:
		return "BC USER: Unknown user error."
	}
}
