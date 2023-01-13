package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	internalError "[MODULE_URL]/pkg/shared/domain/errors"
)

const NO_STATUS_CODE = -1

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()

	err := ctx.Errors.Last()

	if err == nil {
		return
	}

	var appErr *internalError.ErrorApp

	if errors.As(err, &appErr) {
		ctx.JSON(NO_STATUS_CODE, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"errors":  appErr.Fields,
		})

		return
	}

	/*
		var code uint16
		var message string
		var errList []map[string]string
	*/
	/*
		var appErr *errorhandler.ErrorHandler
			var validatorErr validator.ValidationErrors

			//if errors.As(err, &appErr) {
				code = appErr.Code
				message = appErr.Msg
				/*ctx.JSON(NO_STATUS_CODE, gin.H{
					"code":    appErr.Code,
					"message": appErr.Msg,
				})

				return* /
			}

			if errors.As(err, &validatorErr) {
				errList = listOfErrors(validatorErr)

				/*ctx.JSON(NO_STATUS_CODE, gin.H{
					"code":    1,
					"message": "valor del mensaje",
					"errors":  list,
				})

				return* /
		 	}*/

	/*ctx.JSON(NO_STATUS_CODE, gin.H{
		"code":    code,
		"message": message,
		"errors":  errList,
	})

	/*ctx.JSON(NO_STATUS_CODE, gin.H{
		"error": err.Error(),
	})*/
}

/*
func listOfErrors(err validator.ValidationErrors) []map[string]string {
	InvalidFields := make([]map[string]string, 0)

	for _, e := range err {
		errors := map[string]string{}
		errors[e.Field()] = e.Tag()
		InvalidFields = append(InvalidFields, errors)
	}

	return InvalidFields
}

// Este archivo es para hacer un midlaware quizas tiene que llamarse response
/*func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
		resp := map[string]string{"hello": "world"}
		c.JSON(400, resp)

		c.Next()
/**
		c.Next()
		err := c.Errors.Last()

		if err == nil {
			return
		}

		context.AbortWithStatusJSON(statusCode, gin.H{
			"code":    error.Code,
			"message": error.Error(),
		})
		// Use reflect.TypeOf(err.Err) to known the type of your error
		/*	if error, ok := errors.Is(err.Err).(*myspace.KindOfClientError); ok {
			c.JSON(400, gin.H{
				"error": "Blah blahhh",
			})
			return
		}
	}
}

/*
func ErrorHandler(a) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // call c.ext () to execute the middleware
		// Start from here after all middleware and router processing is complete
		// Check for errors in c. elors
		for _, e := range c.Errors {
			err := e.Err
			// Return code and MSG if there is a custom error
			if myErr, ok := err.(Error); ok {
				c.JSON(http.StatusOK, gin.H{
					"code": myErr.Code,
					"msg":  myErr.Msg,
					"data": myErr.Data,
				})
			} else {
				Err.error ()
				// For example, err is set when save session fails
				c.JSON(http.StatusOK, gin.H{
					"code": 500."msg":  "Server exception"."data": err.Error(),
				})
			}
			return // Check for an error
		}
	}
}
/*
type MyError struct {
	Code int
	Msg  string
	Data interface{}}var (
	LOGIN_UNKNOWN = NewError(202."User does not exist")
	LOGIN_ERROR   = NewError(203."Wrong account or password")
	VALID_ERROR   = NewError(300."Parameter error")
	ERROR         = NewError(400."Operation failed")
	UNAUTHORIZED  = NewError(401."You are not logged in.")
	NOT_FOUND     = NewError(404."Resources do not exist")
	INNER_ERROR   = NewError(500."System exception"))func (e *MyError) Error(a) string {
	return e.Msg
}

func NewError(code int, msg string) *MyError {
	return &MyError{
		Msg:  msg,
		Code: code,
	}
}

func GetError(e *MyError, data interface{}) *MyError {
	return &MyError{
		Msg:  e.Msg,
		Code: e.Code,
		Data: data,
	}
}

/*
import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}

		// Use reflect.TypeOf(err.Err) to known the type of your error
		/*	if error, ok := errors.Is(err.Err).(*myspace.KindOfClientError); ok {
			c.JSON(400, gin.H{
				"error": "Blah blahhh",
			})
			return
		}
	}
}
*/
// 200	OK	The request has succeeded.
// 201	Created	The request has succeeded and a new resource has been created as a result of it.
// 202	Accepted	Indicate that the request has been accepted for processing, but the processing has not been completed.
// 204	No Content	The request has succeeded and there is no content to send for this request. This is common for DELETE requests.
// 400	Bad Request	Server cannot process the request due to something wrong with the request.
// 401	Unauthorized	Client needs authentication to get requested resource.
// 403	Forbidden	Client does not have access to the requested resource.
// 404	Not Found	Server cannot find the requested resource.
// 5XX	Server Errors	Server encountered an unexpected condition.

// Tipos de errores
// De validacon
/*
{
    "code": 190,
    "message": "Bad Request",
	"errors" : {
		"first_name": {
			"required": "The First name is required"
		}
	}

}
*/
