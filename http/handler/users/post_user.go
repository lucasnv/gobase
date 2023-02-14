package users

import (
	"github.com/gin-gonic/gin"

	"<MODULE_URL_REPLACE>/http/response"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
	"<MODULE_URL_REPLACE>/pkg/users/application/registeruser"
)

type postUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=4,max=15"`
}

func PostUser(cb commandbus.CommandBus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req postUserRequest

		// Validation request
		if err := ctx.ShouldBind(&req); err != nil {
			response.BadRequest(ctx, err)
			return
		}

		// Command Dispatcher
		_, err := cb.Dispatch(ctx, getPostUserCommand(req))

		if err != nil {
			response.AppError(ctx, err)
			return
		}

		// Return Json response
		response.Success(ctx)
	}
}

func getPostUserCommand(req postUserRequest) commandbus.Command {
	return registeruser.NewCommand(
		req.FirstName,
		req.LastName,
		req.Email,
		req.Password,
	)
}
