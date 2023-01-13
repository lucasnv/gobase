package commandbus

import (
	"[REPO_URL]/internal/commandbus"
	"[REPO_URL]/pkg/shared/insfrastructure/dependencyinjection"
	"[REPO_URL]/pkg/tools/application/createtool"
	"[REPO_URL]/pkg/users/application/registeruser"
)

func Initialize(di *dependencyinjection.DependencyInjection) *commandbus.CommandBus {
	commandBus := commandbus.NewCommandBus()

	commandBus.Register(createtool.COMMMAND_TYPE, di.CreateToolCommandHandler)
	commandBus.Register(registeruser.COMMMAND_TYPE, di.RegisterUserCommandHandler)

	return commandBus
}
