package validators

import (
	"../models"
	"github.com/gin-gonic/gin"
)

// Registration validates user registration struct
type Registration struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// validates user registration object
func (r Registration) validate(c *gin.Context) (models.Model, error) {
	if err := c.ShouldBindJSON(&r); err != nil {
		return models.User{}, err
	}

	user := models.User{
		Email:    r.Email,
		Username: r.Username,
	}

	// hash user's password
	user.HashPassword(r.Password)

	return user, nil
}
