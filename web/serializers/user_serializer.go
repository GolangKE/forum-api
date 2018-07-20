package serializers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

// UserSerializer struct
type UserSerializer struct {
	Context *gin.Context
}

// UserResponse struct
type UserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// Response creates the user response object
func (u *UserSerializer) Response() UserResponse {
	model := u.Context.MustGet("user_model").(models.User)
	user := UserResponse{
		ID:        *(&model.ID),
		Username:  model.Username,
		Email:     model.Email,
		CreatedAt: model.CreatedAt.Unix(),
		UpdatedAt: model.UpdatedAt.Unix(),
	}
	return user
}
