package main

import (
	"os"

	"[MODULE_URL]/http"
	"[MODULE_URL]/internal/config"
	"[MODULE_URL]/pkg/shared/insfrastructure/commandbus"
	"[MODULE_URL]/pkg/shared/insfrastructure/dependencyinjection"
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