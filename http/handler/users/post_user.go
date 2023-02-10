package users

import (
	"github.com/gin-gonic/gin"

	"<MODULE_URL_REPLACE>/http/response"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/commandbus"
	"<MODULE_URL_REPLACE>/pkg/users/application/registeruser"
)

type request struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=4,max=15"`
}

func PostUser(cb commandbus.CommandBus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request

		// Validation request
		if err := ctx.ShouldBind(&req); err != nil {
			response.BadRequest(ctx, err)
			return
		}

		// Command Dispatcher
		if err := cb.Dispatch(ctx, getCommand(req)); err != nil {
			response.AppError(ctx, err)
			return
		}

		// Return Json response
		response.Success(ctx)
	}
}

func getCommand(req request) commandbus.Command {
	return registeruser.NewCommand(
		req.FirstName,
		req.LastName,
		req.Email,
		req.Password,
	)
}
