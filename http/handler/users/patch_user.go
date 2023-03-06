package users

import (
	"github.com/gin-gonic/gin"

	"<MODULE_URL_REPLACE>/http/response"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
	"<MODULE_URL_REPLACE>/pkg/users/application/updateuser"
)

type patchUserRequest struct {
	FirstName string `json:"first_name" binding:""`
	LastName  string `json:"last_name" binding:""`
	Email     string `json:"email" binding:",email"`
}

func PatchUser(cb commandbus.CommandBus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req patchUserRequest

		// Validation request
		if err := ctx.ShouldBind(&req); err != nil {
			response.BadRequest(ctx, err)
			return
		}

		// Command Dispatcher
		_, err := cb.Dispatch(ctx, getPatchUserCommand(req))

		if err != nil {
			response.AppError(ctx, err)
			return
		}

		// Return Json response
		response.Success(ctx)
	}
}

func getPatchUserCommand(req patchUserRequest) commandbus.Command {
	return updateuser.NewCommand(
		req.FirstName,
		req.LastName,
		req.Email,
	)
}
