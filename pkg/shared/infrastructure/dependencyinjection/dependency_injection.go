package dependencyinjection

import (
	"<MODULE_URL_REPLACE>/pkg/users/application/registeruser"
)

type DependencyInjection struct {
	CreateToolCommandHandler   createtool.CommandHandler
	RegisterUserCommandHandler registeruser.CommandHandler
}

func NewDependencyInjection() *DependencyInjection {

	// Repositories

	// Services
	registerUserService := registeruser.NewService()

	// Handlers
	registerUserCommandHandler := registeruser.NewCommandHandler(registerUserService)

	return &DependencyInjection{
		RegisterUserCommandHandler: registerUserCommandHandler,
	}
}
