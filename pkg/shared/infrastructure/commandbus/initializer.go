package commandbus

import (
	"<MODULE_URL_REPLACE>/internal/commandbus"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/dependencyinjection"
	"<MODULE_URL_REPLACE>/pkg/tools/application/createtool"
	"<MODULE_URL_REPLACE>/pkg/users/application/registeruser"
)

func Initialize(di *dependencyinjection.DependencyInjection) *commandbus.CommandBus {
	commandBus := commandbus.NewCommandBus()

	commandBus.Register(createtool.COMMMAND_TYPE, di.CreateToolCommandHandler)
	commandBus.Register(registeruser.COMMMAND_TYPE, di.RegisterUserCommandHandler)

	return commandBus
}
