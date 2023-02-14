//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	
	"<MODULE_URL_REPLACE>/pkg/users"
	"<MODULE_URL_REPLACE>/pkg/users/application/registeruser"
)

func Wire() registeruser.CommandHandler {
	panic(wire.Build(users.ProviderSet))

	return registeruser.CommandHandler{}
}

/*
import (
	"github.com/google/wire"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
	"<MODULE_URL_REPLACE>/pkg/users/application/registeruser"
	"<MODULE_URL_REPLACE>/pkg/users/domain"
	"<MODULE_URL_REPLACE>/pkg/users/infrastructure"
)

type Wire struct {
	// registerUserCommandHandler
}

var WiredRegisterUser = wire.NewSet(
	infrastructure.NewInmemoryUsersRepository,
	registeruser.NewCommandHandler,
	registeruser.NewService,

	wire.Bind(new(commandbus.Handler), new(registeruser.CommandHandler)),
	wire.Bind(new(domain.UserRepository), new(infrastructure.InmemoryUsersRepository)),
)

func setupWire() *Wire {
	wire.Build(
		WiredRegisterUser,

	/*
		// sqldb to DB.
		sqldb.NewConfig,
		sqldb.New,
		wire.Bind(new(DB), new(*sql.DB)),

		// User Repo.
		userrepo.Wired,
		wire.Bind(new(user.Repo), new(*userrepo.Repo)),

		// User Service.
		userservice.Wired,
		wire.Bind(new(user.Service), new(*userservice.Service)),

		// Todo Repo.
		todorepo.Wired,
		wire.Bind(new(todo.Repo), new(*todorepo.Repo)),

		// Todo Service.
		todoservice.Wired,
		wire.Bind(new(todo.Service), new(*todoservice.Service)),

		// This package - the application.
		newApp,

	*
	)

	// Requirement for compilation.
	return &Wire{}
}
*/
