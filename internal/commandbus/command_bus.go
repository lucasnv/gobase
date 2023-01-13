package commandbus

import (
	"context"
)

// Bus defines the expected behaviour from a command bus.
type Bus interface {
	Dispatch(context.Context, Command) error
	Register(Type, Handler)
}

// Type represents an application command type.
type Type string

// Command represents an application command.
type Command interface {
	Type() Type
}

// Handler defines the expected behaviour from a command handler.
type Handler interface {
	Handle(context.Context, Command) error
}
