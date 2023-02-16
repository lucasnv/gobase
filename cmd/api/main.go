package main

import (
	"os"

	"<MODULE_URL_REPLACE>/http"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/commandbus"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/config"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/wire"
	"<MODULE_URL_REPLACE>/pkg/users/application/deleteuser"
	"<MODULE_URL_REPLACE>/pkg/users/application/finduser"
	"<MODULE_URL_REPLACE>/pkg/users/application/registeruser"
)

func main() {

	// CONFIG
	config.InitializeDotEnv()

	// DEPENDENCY INJECTION
	app := wire.Setup()

	// COMMAND BUS
	commandBus := commandbus.NewInMemoryCommandBus()

	registerCommands(commandBus, app)

	// SERVER HTTP
	http.InitializeServer(os.Getenv("SERVER_ADDRESS"), commandBus)
}

// Registar all commands
// TODO: This must be generic, this gonna be use it in many mains
func registerCommands(cb *commandbus.InMemoryCommandBus, app *wire.Wire) {
	cb.Register(registeruser.COMMAND_TYPE, &app.RegisterUserCommandHandler)
	cb.Register(finduser.COMMAND_TYPE, &app.FindUserCommandHandler)
	cb.Register(deleteuser.COMMAND_TYPE, &app.DeleteUserCommandHandler)
}
