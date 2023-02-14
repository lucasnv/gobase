package main

import (
	"os"

	"github.com/lucasnv/gobase-test/http"
	"github.com/lucasnv/gobase-test/pkg/shared/infrastructure/commandbus"
	"github.com/lucasnv/gobase-test/pkg/shared/infrastructure/config"
	"github.com/lucasnv/gobase-test/pkg/shared/infrastructure/wire"
	"github.com/lucasnv/gobase-test/pkg/users/application/finduser"
	"github.com/lucasnv/gobase-test/pkg/users/application/registeruser"
)

func main() {

	// CONFIG
	config.InitializeDotEnv()

	// DEPENDENCY INJECTION
	app := wire.Setup()

	// COMMAND BUS
	commandBus := commandbus.NewInMemoryCommandBus()
	//TODO: Generar un funcion en otro archivo que registre todo esto
	commandBus.Register(registeruser.COMMAND_TYPE, &app.RegisterUserCommandHandler)
	commandBus.Register(finduser.COMMAND_TYPE, &app.FindUserCommandHandler)

	// SERVER HTTP
	http.InitializeServer(os.Getenv("SERVER_ADDRESS"), commandBus)
}
