package users

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"<MODULE_URL_REPLACE>/http/response"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
	"<MODULE_URL_REPLACE>/pkg/users/application/finduser"
)

type getUserRequest struct {
	Id string `json:"first_name" binding:"required"`
}

func GetUser(cb commandbus.CommandBus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req getUserRequest

		// Validation request
		if err := ctx.ShouldBind(&req); err != nil {
			response.BadRequest(ctx, err)
			return
		}

		// Command Dispatcher
		if user, err := cb.Dispatch(ctx, getGetUserQuery(req)); err != nil {
			response.AppError(ctx, err)
			return
		} else {
			fmt.Printf("%v\n", user)
		}

		// Return Json response
		response.Success(ctx)
	}
}

func getGetUserQuery(req getUserRequest) commandbus.Command {
	return finduser.NewCommand(
		req.Id,
	)
}
