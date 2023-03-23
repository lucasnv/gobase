package commandbus

import (
	"context"

	"github.com/gin-gonic/gin"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
)

// Type represents an application command type.
type Type string

// Response represent a struct with the information
type Response any

type Handlers map[Type]Handler

type Middleware func(next HandlerFunc) HandlerFunc

type HandlerFunc func(ctx *context.Context, cmd Command) (Response, errors.App)

// Command represents an application command.
type Command interface {
	Type() Type
}

// Bus defines the expected behaviour from a command bus.
type CommandBus interface {
	Dispatch(ctx *gin.Context, cmd Command) (Response, errors.App)
	RegisterCommand(t Type, h Handler)
	RegisterMiddleware(m Middleware)
}

// Handler defines the expected behaviour from a command handler.
type Handler interface {
	Handle(ctx *context.Context, cmd Command) (Response, errors.App)
}
