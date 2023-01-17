package main

import (
	"os"

	"<MODULE_URL_REPLACE>/http"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/config"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/commandbus"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/dependencyinjection"
)

func main() {
	// CONFIG
	config.InitializeDotEnv()

	// DEPENDENCY INJECTION
	dependencyInjection := dependencyinjection.NewDependencyInjection()

	// COMMAND BUS
	commandBus := commandbus.Initialize(dependencyInjection)

	// SERVER HTTP
	http.InitializeServer(os.Getenv("SERVER_ADDRESS"), commandBus)
}
