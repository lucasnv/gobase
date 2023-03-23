package main

import (
	"os"

	"<MODULE_URL_REPLACE>/http"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/commandbus"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/commandbus/middleware"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/config"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/storage"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/wire"
	"<MODULE_URL_REPLACE>/pkg/users/application/deleteuser"
	"<MODULE_URL_REPLACE>/pkg/users/application/finduser"
	"<MODULE_URL_REPLACE>/pkg/users/application/findusers"
	"<MODULE_URL_REPLACE>/pkg/users/application/registeruser"
	"<MODULE_URL_REPLACE>/pkg/users/application/updateuser"
)

func main() {

	// CONFIG
	config.InitializeDotEnv()

	// MONGO INITIALIZATION
	storage.InitializeMongoConnection()

	// DEPENDENCY INJECTION
	app := wire.Setup()

	// COMMAND BUS
	commandBus := commandbus.NewInMemoryCommandBus()

	registerCommands(commandBus, app)

	// SERVER HTTP
	http.InitializeServer(os.Getenv("SERVER_ADDRESS"), commandBus)
}

// Configure commands and middleware
func registerCommands(bus *commandbus.InMemoryCommandBus, app *wire.Wire) {
	bus.RegisterCommand(registeruser.COMMAND_TYPE, &app.RegisterUserCommandHandler)
	bus.RegisterCommand(finduser.COMMAND_TYPE, &app.FindUserCommandHandler)
	bus.RegisterCommand(findusers.COMMAND_TYPE, &app.FindUsersCommandHandler)
	bus.RegisterCommand(deleteuser.COMMAND_TYPE, &app.DeleteUserCommandHandler)
	bus.RegisterCommand(updateuser.COMMAND_TYPE, &app.UpdateUserCommandHandler)

	bus.RegisterMiddleware(middleware.Transaction)
	bus.RegisterMiddleware(middleware.Event)
}
