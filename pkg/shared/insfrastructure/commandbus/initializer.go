package commandbus

import (
	"[MODULE_URL]/internal/commandbus"
	"[MODULE_URL]/pkg/shared/insfrastructure/dependencyinjection"
	"[MODULE_URL]/pkg/tools/application/createtool"
	"[MODULE_URL]/pkg/users/application/registeruser"
)

func Initialize(di *dependencyinjection.DependencyInjection) *commandbus.CommandBus {
	commandBus := commandbus.NewCommandBus()

	commandBus.Register(createtool.COMMMAND_TYPE, di.CreateToolCommandHandler)
	commandBus.Register(registeruser.COMMMAND_TYPE, di.RegisterUserCommandHandler)

	return commandBus
}
