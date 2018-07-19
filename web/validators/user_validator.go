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

	user := models.User{}
	user.Email = r.Email
	user.Username = r.Username
	user.PasswordDigest = r.Password

	return user, nil
}
