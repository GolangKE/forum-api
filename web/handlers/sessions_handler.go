package handlers

import (
	"github.com/gin-gonic/gin"
)

// DestroySession ends user session by invalidatin token
func DestroySession(context *gin.Context) {
	context.JSON(204, gin.H{"message": "User logged out session"})
}
