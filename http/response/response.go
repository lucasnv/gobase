package response

import (
	"github.com/gin-gonic/gin"
)

func WithSuccess(ctx *gin.Context, statusCode int) {
	ctx.Status(statusCode)
}

func WithError(ctx *gin.Context, statusCode int, err error) {
	ctx.Status(statusCode)
	ctx.Error(err)
	ctx.Abort()
}

// Response type
/*

{
	"data": {
		"first_name": "Lucas",
		"last_name": "Vazquez",
		"email": "lucasnvazquez@gmail.com"
		"created_at": "2020-12-10 12:34:56 UTC"
	}
}
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
{
	"code": 1001
	"message" : "El usuario no existe"
}

{
	"code": 10
	"message" : "Invalid fields"
	"errors": [
		"first_name": "required",
		"last_name": "required",
	]
}
*/