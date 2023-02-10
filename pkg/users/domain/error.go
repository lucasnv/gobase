package domain

import (
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/validator"
)

// Users' errors goes from 1000 to 1999
const (
	UNKNOWN_USER_ERROR = 1000
	INVALID_USER       = 1001
)

func NewUserError(c errors.ErrCode, ve ...validator.ValidationError) *errors.AppError {
	var error errors.AppError = errors.AppError{
		ErrCode:    c,
		ErrMessage: getErrorMessage(c),
	}

	if ve != nil {
		for i, v := range ve {
			error.AddError(v.Field(), v.Error())
		}
	}

	return &error
}

func getErrorMessage(c errors.ErrCode) string {
	switch c {
	case INVALID_USER:
		return "BC USER: cannot process the request due to something wrong with the data."
	case UNKNOWN_USER_ERROR:
		return "BC USER: Unknown user error."
	default:
		return "BC USER: Unknown user error."
	}
}
