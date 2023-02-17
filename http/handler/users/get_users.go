package users

import (
	"github.com/gin-gonic/gin"

	"<MODULE_URL_REPLACE>/http/request"
	"<MODULE_URL_REPLACE>/http/response"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
	"<MODULE_URL_REPLACE>/pkg/users/application/findusers"
)

func GetUsers(cb commandbus.CommandBus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request.FilterRequest

		// Validation request
		if err := ctx.ShouldBindQuery(&req); err != nil {
			response.BadRequest(ctx, err)
			return
		}

		// Command Dispatcher
		user, err := cb.Dispatch(ctx, getGetUsersCommand(req))

		if err != nil {
			response.AppError(ctx, err)
			return
		}

		// Return Json response
		response.SuccessWithData(ctx, user)
	}
}

func getGetUsersCommand(req request.FilterRequest) commandbus.Command {
	return findusers.NewCommand(
		req.Filter,
		req.SortBy,
		req.SortOrder,
		req.Page,
		req.PerPage,
	)
}
