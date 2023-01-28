//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"<MODULE_URL_REPLACE>/pkg/users/application/registeruser"
)

func Wire() registeruser.CommandHandler {
	panic(wire.Build(ProviderSet))

	return registeruser.CommandHandler{}
}
