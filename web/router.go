package web

import (
	"github.com/gin-gonic/gin"
)

// Register the app's router, accepts an instance
// of engine with default middleware already attached
func Register(router *gin.Engine) {
	// group API routes
	api := router.Group("/api/v1")

	api.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "pong"})
	})

}
