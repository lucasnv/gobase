package commandbus

import (
	"context"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
)

// Type represents an application command type.
type Type string

// Bus defines the expected behaviour from a command bus.
type CommandBus interface {
	Dispatch(context.Context, Command) errors.App
	Register(Type, Handler)
}

// Command represents an application command.
type Command interface {
	Type() Type
}

// Handler defines the expected behaviour from a command handler.
type Handler interface {
	Handle(context.Context, Command) errors.App
}
