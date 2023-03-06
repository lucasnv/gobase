package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	v10 "github.com/go-playground/validator/v10"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/collection"
	appError "<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
)

//	GET: 200 OK
//	POST: 201 Created
//	PUT: 200 OK
//	PATCH: 200 OK
//	DELETE: 204 No Content

// 200	OK	The request has succeeded.
// 201	Created	The request has succeeded and a new resource has been created as a result of it.
// 202	Accepted	Indicate that the request has been accepted for processing, but the processing has not been completed.
// 204	No Content	The request has succeeded and there is no content to send for this request. This is common for DELETE requests.
// 400	Bad Request	Server cannot process the request due to something wrong with the request.
// 401	Unauthorized	Client needs authentication to get requested resource.
// 403	Forbidden	Client does not have access to the requested resource.
// 404	Not Found	Server cannot find the requested resource.
// 5XX	Server Errors	Server encountered an unexpected condition.

/**********************************
> Success response type examples <
***********************************/

//           Single Data
// *******************************
//	{
//		"data": {
//			"first_name": "Lucas",
//			"last_name": "Vazquez",
//			"email": "lucasnvazquez@gmail.com"
//			"created_at": "2020-12-10 12:34:56 UTC"
//		}
//	}

//              List
// *******************************
//	{
//		"data": [
//			{
//				"first_name": "Lucas",
//				"last_name": "Vazquez",
//				"email": "lucasnvazquez@gmail.com"
//				"created_at": "2020-12-10 12:34:56 UTC"
//			},
//			{
//				"first_name": "Jorge",
//				"last_name": "Caruzo",
//				"email": "jorgecaruzo@gmail.com"
//				"created_at": "2021-10-11 12:34:56 UTC"
//			}
//		],
//      "metadata": {
//
//		}
//	}

/***********************************
>   Error response type examples   <
************************************/

//         Single error
// *******************************
//	{
//		"code": 1001
//		"message" : "El usuario no existe"
//	}

//        Multi fields error
// *******************************
//	{
//		"code": 10
//		"message" : "Invalid fields"
//		"errors": [
//			"first_name": "required",
//			"last_name": "required",
//		]
//	}

func Success(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func Created(ctx *gin.Context) {
	ctx.Status(http.StatusCreated)
}

func NoContent(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func SuccessWithData(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func SuccessWithCollection(ctx *gin.Context, data any) {
	c := data.(collection.Collection)
	ctx.JSON(http.StatusOK, gin.H{
		"data":     c.Data(),
		"metadata": c.Metadata(),
	})
}

func GenericError(ctx *gin.Context, err error) {
	var appErr = appError.NewAppError(appError.UNKNOWN_ERROR)
	ctx.Error(appErr)
}

// TODO: I have to delete the v10 Lib dependency implementing my validation
func BadRequest(ctx *gin.Context, err error) {
	var validatorErr v10.ValidationErrors
	var appErr *appError.AppError = appError.NewAppError(appError.BAD_REQUEST_ERROR)

	if errors.As(err, &validatorErr) {
		for _, field := range validatorErr {
			appErr.AddError(field.Field(), field.Tag())
		}
	}

	ctx.Error(appErr)
}

func AppError(ctx *gin.Context, err appError.App) {
	ctx.Error(err)
}
