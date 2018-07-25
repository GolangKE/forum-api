package middlewares

import (
	"net/http"
	"time"

	"../models"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func AuthMiddleware() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:         "golangke.org",
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: authenticator,
		Authorizator:  authorizator,
		Unauthorized:  unauthorized,
		TokenLookup:   "header: Authorization, query: token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		LoginResponse: loginResponse,
		PayloadFunc:   payload,
	}
}

func authenticator(userID string, password string, c *gin.Context) (interface{}, bool) {
	condition := map[string]interface{}{"email": userID}

	if data, err := models.FindOne(models.User{}, condition); err == nil {
		user, _ := data.(models.User)

		if err := user.CheckPassword(password); err == nil {
			return user, true
		}
	}

	return nil, false
}

func authorizator(user interface{}, c *gin.Context) bool {
	// TODO: Implement role-based authorization here
	return true
}

func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func loginResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"token":  token,
		"expire": expire.Unix(),
	})
}

func payload(data interface{}) jwt.MapClaims {
	user := data.(models.User)

	return map[string]interface{}{
		"id":       *(&user.ID),
		"email":    user.Email,
		"username": user.Username,
		"role":     "member",
	}
}
