package middlewares

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"../models"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:            "golangke.org",
		Key:              []byte("secret key"),
		SigningAlgorithm: "HS256",
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		Authenticator:    authenticator,
		Authorizator:     authorizator,
		Unauthorized:     unauthorized,
		TokenLookup:      "header: Authorization, query: token",
		TokenHeadName:    "Bearer",
		TimeFunc:         time.Now,
		LoginResponse:    loginResponse,
		RefreshResponse:  refreshResponse,
		PayloadFunc:      payload,
	}
}

func authenticator(userID string, password string, c *gin.Context) (interface{}, bool) {
	condition := map[string]interface{}{"email": userID}

	if data, err := models.FindOne(models.User{}, condition); err == nil {
		user := data.(models.User)

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
	// Extract claims from token
	claims := claimsFromToken(token)

	// save token to db and return response
	models.Create(models.UserToken{
		Token:     token,
		ExpiresAt: expire,
		UserID:    claims["id"].(string),
		IsValid:   true,
	})

	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"token":  token,
		"expire": expire.Unix(),
	})
}

func refreshResponse(c *gin.Context, code int, token string, expire time.Time) {
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

func claimsFromToken(token string) map[string]interface{} {
	var claims map[string]interface{}
	payload := strings.Split(token, ".")[1]

	// add padding to yield appropriate decoded data
	// https://github.com/golang/go/issues/4237
	if length := len(payload) % 4; length > 0 {
		payload += strings.Repeat("=", 4-length)
	}

	data, _ := base64.URLEncoding.DecodeString(payload)
	json.Unmarshal(data, &claims)

	return claims
}
