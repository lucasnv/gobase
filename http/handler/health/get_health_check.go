package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealthCheckGet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "healthy API status.",
		})
	}
}
