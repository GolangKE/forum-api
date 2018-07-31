package handlers

import (
	"../models"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// DestroySession ends user session by invalidatin token
func DestroySession(context *gin.Context) {
	token := models.UserToken{Token: jwt.GetToken(context)}
	data, _ := token.FindOne()

	data.(*models.UserToken).InvalidateToken()
	context.JSON(204, gin.H{"message": "logged out"})
}
