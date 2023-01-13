package http

import (
	"<MODULE_URL_REPLACE>/http/handler/health"
)

//TODO: Change this constant version to GIN GROUP ROUTE
const VERSION = "v1"

// ConfigureRoutes Configure each route with his own handler
func ConfigureRoutes(server Server) {

	// HEALTH CHECK
	server.engine.GET("/"+VERSION+"/health-check", health.GetHealthCheckGet())

	// ADD CUSTOM HANDLERS
	//server.engine.POST("/"+VERSION+"/users", users.PostUser(server.commandBus))
}
