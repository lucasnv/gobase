package users

import (
	"github.com/gin-gonic/gin"

	"<MODULE_URL_REPLACE>/http/response"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
	"<MODULE_URL_REPLACE>/pkg/users/application/deleteuser"
)

type deleteUserRequest struct {
	Id string `uri:"id" binding:"required"`
}

func DeleteUser(cb commandbus.CommandBus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req deleteUserRequest

		// Validation request
		if err := ctx.ShouldBindUri(&req); err != nil {
			response.BadRequest(ctx, err)
			return
		}

		// Command Dispatcher
		_, err := cb.Dispatch(ctx, getDeleteUserCommand(req))

		if err != nil {
			response.AppError(ctx, err)
			return
		}

		// Return Json response
		response.NoContent(ctx)
	}
}

func getDeleteUserCommand(req deleteUserRequest) commandbus.Command {
	return deleteuser.NewCommand(
		req.Id,
	)
}
