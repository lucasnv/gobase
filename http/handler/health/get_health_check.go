package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHealthCheckGet Health Check Handler
func GetHealthCheckGet() gin.HandlerFunc {
	return func(context *gin.Context) {
		msg := "healthy API status."
		context.JSON(http.StatusOK, gin.H{
			"status": msg,
		})
	}
}
