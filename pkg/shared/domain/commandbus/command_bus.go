package commandbus

import (
	"context"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
)

// Type represents an application command type.
type Type string

// Command represents an application command.
type Command interface {
	Type() Type
}

// Response represent a struct with the information
type Reponse interface {
}

// Bus defines the expected behaviour from a command bus.
type CommandBus interface {
	Dispatch(context.Context, Command) (Reponse, errors.App)
	Register(Type, Handler)
}

// Handler defines the expected behaviour from a command handler.
type Handler interface {
	Handle(context.Context, Command) (Reponse, errors.App)
}
