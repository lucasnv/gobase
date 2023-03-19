package users

import (
	"github.com/gin-gonic/gin"

	"<MODULE_URL_REPLACE>/http/response"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
	"<MODULE_URL_REPLACE>/pkg/users/application/updateuser"
)

type putUserIdRequest struct {
	Id string `uri:"id" binding:"required"`
}
type putUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

func PutUser(cb commandbus.CommandBus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqId putUserIdRequest
		var req putUserRequest

		// Validation request
		if err := ctx.ShouldBindUri(&reqId); err != nil {
			response.BadRequest(ctx, err)
			return
		}

		if err := ctx.ShouldBind(&req); err != nil {
			response.BadRequest(ctx, err)
			return
		}

		// Command Dispatcher
		_, err := cb.Dispatch(ctx, getPutUserCommand(reqId, req))

		if err != nil {
			response.AppError(ctx, err)
			return
		}

		// Return Json response
		response.Success(ctx)
	}
}

func getPutUserCommand(reqId putUserIdRequest, req putUserRequest) commandbus.Command {
	return updateuser.NewCommand(
		reqId.Id,
		req.FirstName,
		req.LastName,
		req.Email,
	)
}
