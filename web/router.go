package web

import (
	"./handlers"
	"./middlewares"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// Register the app's router, accepts an instance
// of engine with default middleware already attached
func Register(router *gin.Engine) {
	// Register App's middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// group API routes
	api := router.Group("/api/v1")

	// ping check route
	api.GET("/ping", authMiddleware().MiddlewareFunc(), handlers.Ping)

	// Group Users API endpoints
	usersEndpoint(api.Group("/users"))

	// Group Sessions API endpoints
	sessionsEndpoint(api.Group("/sessions"))
}

func usersEndpoint(router *gin.RouterGroup) {
	router.POST("/new", handlers.NewUser)
}

func sessionsEndpoint(router *gin.RouterGroup) {
	router.POST("/new", authMiddleware().LoginHandler)

	authRequired := router.Group("/", authMiddleware().MiddlewareFunc())
	{
		authRequired.GET("/refresh_token", authMiddleware().RefreshHandler)
		authRequired.DELETE("/", handlers.DestroySession)
	}
}

func authMiddleware() *jwt.GinJWTMiddleware {
	return middlewares.AuthMiddleware()
}
