package commandbus

import (
	"context"

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

func (cb *InMemoryCommandBus) Dispatch(ctx context.Context, cmd commandbus.Command) (commandbus.Reponse, errors.App) {
	handler, ok := cb.handlers[cmd.Type()]

	if !ok {
		return nil, errors.NewAppError(errors.UNEXPECTED_COMMAND_ERROR)
	}

	return handler.Handle(ctx, cmd)
}

func (b *InMemoryCommandBus) Register(t commandbus.Type, h commandbus.Handler) {
	b.handlers[t] = h
}

var _ commandbus.CommandBus = (*InMemoryCommandBus)(nil)
