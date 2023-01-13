package main

import (
	"os"

	"[REPO_URL]/http"
	"[REPO_URL]/internal/config"
	"[REPO_URL]/pkg/shared/insfrastructure/commandbus"
	"[REPO_URL]/pkg/shared/insfrastructure/dependencyinjection"
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
