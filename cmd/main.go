package main

import (
	"os"

	"<MODULE_URL_REPLACE>/cmd/di"
	"<MODULE_URL_REPLACE>/http"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/commandbus"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/config"

	"<MODULE_URL_REPLACE>/pkg/users/application/registeruser"
)

func main() {

	// CONFIG
	config.InitializeDotEnv()

	// DEPENDENCY INJECTION
	// Handlers
	registerUserCommandHandler := di.Wire()
	//registeruser.NewCommandHandler(registerUserService)

	// COMMAND BUS
	commandBus := commandbus.NewInMemoryCommandBus()

	commandBus.Register(registeruser.COMMMAND_TYPE, registerUserCommandHandler)

	// SERVER HTTP
	http.InitializeServer(os.Getenv("SERVER_ADDRESS"), commandBus)
}
