package commandbus

import (
	"context"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
)

// CommandBus is an in-memory implementation of the command.Bus.
type InMemoryCommandBus struct {
	handlers map[commandbus.Type]commandbus.Handler
}

// NewInMemoryCommandBus initializes a new instance of CommandBus.
func NewInMemoryCommandBus() *InMemoryCommandBus {
	return &InMemoryCommandBus{
		handlers: make(map[commandbus.Type]commandbus.Handler),
	}
}

// Dispatch implements the command.Bus interface.
func (cb *InMemoryCommandBus) Dispatch(ctx context.Context, cmd commandbus.Command) (commandbus.Reponse, errors.App) {
	handler, ok := cb.handlers[cmd.Type()]

	if !ok {
		return nil, nil // TODO: Generate an app error when a command does not exist, check if that if is the real problem
	}

	return handler.Handle(ctx, cmd)
}

// Register implements the command.Bus interface.
func (b *InMemoryCommandBus) Register(t commandbus.Type, h commandbus.Handler) {
	b.handlers[t] = h
}

var _ commandbus.CommandBus = (*InMemoryCommandBus)(nil)
