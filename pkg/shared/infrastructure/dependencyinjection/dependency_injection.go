package dependencyinjection

import (
	"<MODULE_URL_REPLACE>/pkg/tools/application/createtool"
	"<MODULE_URL_REPLACE>/pkg/users/application/registeruser"
)

type DependencyInjection struct {
	CreateToolCommandHandler   createtool.CommandHandler
	RegisterUserCommandHandler registeruser.CommandHandler
}

func NewDependencyInjection() *DependencyInjection {

	// Repositories

	// Services
	createToolService := createtool.NewService()
	registerUserService := registeruser.NewService()

	// Handlers
	createToolCommandHandler := createtool.NewCommandHandler(createToolService)
	registerUserCommandHandler := registeruser.NewCommandHandler(registerUserService)

	return &DependencyInjection{
		CreateToolCommandHandler:   createToolCommandHandler,
		RegisterUserCommandHandler: registerUserCommandHandler,
	}
}
