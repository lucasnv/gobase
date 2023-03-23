package commandbus

import (
	"github.com/gin-gonic/gin"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
)

// CommandBus is an in-memory implementation of the command.Bus.
type InMemoryCommandBus struct {
	handlers    commandbus.Handlers
	middlewares []commandbus.Middleware
}

// NewInMemoryCommandBus initializes a new instance of CommandBus.
func NewInMemoryCommandBus() *InMemoryCommandBus {
	return &InMemoryCommandBus{
		handlers:    make(commandbus.Handlers),
		middlewares: make([]commandbus.Middleware, 0),
	}
}

func (bus *InMemoryCommandBus) Dispatch(ctx *gin.Context, cmd commandbus.Command) (commandbus.Response, errors.App) {
	handler, ok := bus.handlers[cmd.Type()]

	if !ok {
		return nil, errors.NewAppError(errors.UNEXPECTED_COMMAND_ERROR)
	}

	var finalHandler commandbus.HandlerFunc = handler.Handle

	for i := len(bus.middlewares) - 1; i >= 0; i-- {
		finalHandler = bus.middlewares[i](finalHandler)
	}

	c := ctx.Request.Context()

	return finalHandler(&c, cmd)
}

func (b *InMemoryCommandBus) RegisterCommand(t commandbus.Type, h commandbus.Handler) {
	b.handlers[t] = h
}

func (bus *InMemoryCommandBus) RegisterMiddleware(m commandbus.Middleware) {
	bus.middlewares = append(bus.middlewares, m)
}

var _ commandbus.CommandBus = (*InMemoryCommandBus)(nil)
