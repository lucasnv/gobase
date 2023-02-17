package http

import (
	"<MODULE_URL_REPLACE>/http/handler/health"
	"<MODULE_URL_REPLACE>/http/handler/users"
)

// ConfigureRoutes Configure each route with his own handler
func ConfigureRoutes(server Server) {

	v1 := server.engine.Group("/v1")
	{
		// HEALTH CHECK
		v1.GET("/health-check", health.GetHealthCheckGet())

		// USERS HANDLERS
		v1.POST("/users", users.PostUser(server.commandBus))
		v1.GET("/users/:id", users.GetUser(server.commandBus))
		v1.GET("/users", users.GetUsers(server.commandBus))
		v1.DELETE("/users/:id", users.DeleteUser(server.commandBus))
	}

	// ADD CUSTOM HANDLERS
}
