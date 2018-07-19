package serializers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

type UserSerializer struct {
	Context *gin.Context
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (u *UserSerializer) Response() UserResponse {
	model := u.Context.MustGet("user_model").(models.User)
	user := UserResponse{
		ID:       *(&model.ID),
		Username: model.Username,
		Email:    model.Email,
	}
	return user
}
