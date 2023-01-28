package commandbus

import (
	"context"
)

// CommandBus is an in-memory implementation of the command.Bus.
type InMemoryCommandBus struct {
	handlers map[Type]Handler
}

// NewInMemoryCommandBus initializes a new instance of CommandBus.
func NewInMemoryCommandBus() *InMemoryCommandBus {
	return &InMemoryCommandBus{
		handlers: make(map[Type]Handler),
	}
}

// Dispatch implements the command.Bus interface.
func (cb *InMemoryCommandBus) Dispatch(ctx context.Context, cmd Command) error {
	handler, ok := cb.handlers[cmd.Type()]

	if !ok {
		return nil
	}

	return handler.Handle(ctx, cmd)
}

// Register implements the command.Bus interface.
func (b *InMemoryCommandBus) Register(t Type, h Handler) {
	b.handlers[t] = h
}

var _ CommandBus = (*InMemoryCommandBus)(nil)
