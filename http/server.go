package http

import (
	"log"

	"github.com/gin-gonic/gin"

	"<MODULE_URL_REPLACE>/http/middleware"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/commandbus"
)

type Server struct {
	engine     *gin.Engine
	commandBus commandbus.CommandBus
}

// InitializeServer Init api server
func InitializeServer(address string, cm commandbus.CommandBus) {
	engine := gin.Default()

	engine.Use(middleware.ErrorHandler())

	server := Server{
		engine:     engine,
		commandBus: cm,
	}

	ConfigureRoutes(server)
	runServer(server, address)
}

func runServer(server Server, address string) {
	err := server.engine.Run(address)

	if err != nil {
		log.Fatal(err)
	}
}
