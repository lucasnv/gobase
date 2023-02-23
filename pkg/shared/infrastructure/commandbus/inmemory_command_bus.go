package commandbus

import (
	"github.com/gin-gonic/gin"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
)

// CommandBus is an in-memory implementation of the command.Bus.
type InMemoryCommandBus struct {
	handlers commandbus.Handlers
}

// NewInMemoryCommandBus initializes a new instance of CommandBus.
func NewInMemoryCommandBus() *InMemoryCommandBus {
	return &InMemoryCommandBus{
		handlers: make(commandbus.Handlers),
	}
}

func (cb *InMemoryCommandBus) Dispatch(ctx *gin.Context, cmd commandbus.Command) (commandbus.Response, errors.App) {
	handler, ok := cb.handlers[cmd.Type()]

	c := ctx.Request.Context()

	if !ok {
		return nil, errors.NewAppError(errors.UNEXPECTED_COMMAND_ERROR)
	}

	return handler.Handle(&c, cmd)
}

func (b *InMemoryCommandBus) Register(t commandbus.Type, h commandbus.Handler) {
	b.handlers[t] = h
}

var _ commandbus.CommandBus = (*InMemoryCommandBus)(nil)
