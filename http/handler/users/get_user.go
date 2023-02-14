package users

import (
	"github.com/gin-gonic/gin"

	"<MODULE_URL_REPLACE>/http/response"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
	"<MODULE_URL_REPLACE>/pkg/users/application/finduser"
)

type getUserRequest struct {
	Id string `uri:"id" binding:"required"`
}

func GetUser(cb commandbus.CommandBus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req getUserRequest

		// Validation request
		if err := ctx.ShouldBindUri(&req); err != nil {
			response.BadRequest(ctx, err)
			return
		}

		// Command Dispatcher
		user, err := cb.Dispatch(ctx, getGetUserCommand(req))

		if err != nil {
			response.AppError(ctx, err)
			return
		}

		// Return Json response
		response.SuccessWithData(ctx, user)
	}
}

func getGetUserCommand(req getUserRequest) commandbus.Command {
	return finduser.NewCommand(
		req.Id,
	)
}
