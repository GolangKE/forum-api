package handlers

import (
	"net/http"

	"../models"
	"../serializers"
	"../validators"
	"github.com/gin-gonic/gin"
)

// NewUser handles new user registration requests
// from api/v1/users/new
func NewUser(context *gin.Context) {
	validator := validators.Registration{}

	if user, err := validators.Validate(validator, context); err != nil {
		context.JSON(http.StatusUnprocessableEntity, validators.Error(err))
	} else {
		if user, err := models.Create(user); err != nil {
			context.JSON(http.StatusUnprocessableEntity, models.Error(err))
		} else {
			context.Set("user_model", user)
			serializer := serializers.UserSerializer{context}
			context.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
		}
	}
}
