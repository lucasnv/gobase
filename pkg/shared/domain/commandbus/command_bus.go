package commandbus

import (
	"context"

	"github.com/gin-gonic/gin"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
)

// Type represents an application command type.
type Type string

type Handlers map[Type]Handler

// Command represents an application command.
type Command interface {
	Type() Type
}

// Response represent a struct with the information
type Response interface {
}

// Bus defines the expected behaviour from a command bus.
type CommandBus interface {
	Dispatch(ctx *gin.Context, c Command) (Response, errors.App)
	Register(t Type, h Handler)
}

// Handler defines the expected behaviour from a command handler.
type Handler interface {
	Handle(ctx *context.Context, c Command) (Response, errors.App)
}
