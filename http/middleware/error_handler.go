package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	appError "<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	userDomain "<MODULE_URL_REPLACE>/pkg/users/domain"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, e := range c.Errors {
			ctxErr := e.Err

			if err, ok := ctxErr.(appError.App); ok {
				c.JSON(getHttpCodeByAppErr(err.Code()), gin.H{
					"code":    err.Code(),
					"message": err.Message(),
					"errors":  err.Errors(),
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": "Internal Server error",
					"errors":  err.Error(),
				})
			}

			return
		}
	}
}

// getHttpCodeByAppErr returns a HTTP status code. returns a 501 if there is not any match
func getHttpCodeByAppErr(code appError.ErrCode) int {
	switch code {

	// Mapping generic app errors
	case appError.BAD_REQUEST_CODE:
		return http.StatusBadRequest
	case appError.UNAUTHORIZED_CODE:
		return http.StatusUnauthorized
	case appError.FORBIDDEN_CODE:
		return http.StatusForbidden
	case appError.NOT_FOUND_CODE:
		return http.StatusNotFound
	case appError.INNER_ERROR_CODE:
		return http.StatusInternalServerError

	// Mapping user bc errors
	case userDomain.INVALID_USER_ERROR:
		return http.StatusBadRequest

	// Default response
	default:
		return http.StatusNotImplemented
	}
}

// TODO: I have to map all errors
