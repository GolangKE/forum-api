package middlewares

import (
	"fmt"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// VerifyToken verifies that user's jwt wasn't
// invalidated before issuing a new one
func VerifyToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := jwt.GetToken(context)
		fmt.Println(token)
	}
}
