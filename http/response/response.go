package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	GET: 200 OK
	POST: 201 Created
	PUT: 200 OK
	PATCH: 200 OK
	DELETE: 204 No Content
*/

/*
	Success response type examples

- Single Data -

	{
		"data": {
			"first_name": "Lucas",
			"last_name": "Vazquez",
			"email": "lucasnvazquez@gmail.com"
			"created_at": "2020-12-10 12:34:56 UTC"
		}
	}

- List -

	{
		"data": [
			{
				"first_name": "Lucas",
				"last_name": "Vazquez",
				"email": "lucasnvazquez@gmail.com"
				"created_at": "2020-12-10 12:34:56 UTC"
			},
			{
				"first_name": "Jorge",
				"last_name": "Caruzo",
				"email": "jorgecaruzo@gmail.com"
				"created_at": "2021-10-11 12:34:56 UTC"
			}
		]
	}
*/
func WithSuccess(ctx *gin.Context, statusCode int) {
	ctx.Status(statusCode)
}

/* Error response type examples
- Single error -
{
	"code": 1001
	"message" : "El usuario no existe"
}

- Multi fields error -
{
	"code": 10
	"message" : "Invalid fields"
	"errors": [
		"first_name": "required",
		"last_name": "required",
	]
}
*/

type responseCode struct {
	code    int
	message string
}

// Response codes
var badRequest responseCode = responseCode{code: 10, message: "Invalid fields"}

func Success(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func Created(ctx *gin.Context) {
	ctx.Status(http.StatusCreated)
}

func BadRequest(ctx *gin.Context, err error) {

	/* Example
	{
		"code": 10
		"message" : "Invalid fields"
		"errors": [
			"first_name": "required",
			"last_name": "required",
		]
	}
	*/

	/*fmt.Println(err)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code":badRequest.code, "message": badRequest.message, "errors": err})
	*/

	appErr := BAD_REQUEST
	appErr.Errors = err.Error()
	ctx.Error(appErr)
}
