package handlers

import (
	"github.com/gin-gonic/gin"
)

// Ping check
func Ping(context *gin.Context) {
	context.JSON(200, gin.H{"message": "pong"})
}
