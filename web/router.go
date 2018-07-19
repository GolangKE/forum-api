package web

import (
	"./handlers"
	"github.com/gin-gonic/gin"
)

// Register the app's router, accepts an instance
// of engine with default middleware already attached
func Register(router *gin.Engine) {
	// group API routes
	api := router.Group("/api/v1")

	// ping check route
	api.GET("/ping", handlers.Ping)

	// new user with email/password route
	api.POST("/users/new", handlers.NewUser)
}
