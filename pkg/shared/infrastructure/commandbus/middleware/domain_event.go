package middleware

import (
	"context"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
)

var Event commandbus.Middleware = func(next commandbus.HandlerFunc) commandbus.HandlerFunc {

	return func(ctx *context.Context, cmd commandbus.Command) (commandbus.Response, errors.App) {
		//defer log.Printf("Event after command")

		return next(ctx, cmd)
	}
}
