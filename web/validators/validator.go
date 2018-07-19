package validators

import (
	"fmt"

	"../models"
	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v8"
)

// Validator is forum-api's main validator interface
type Validator interface {
	validate(c *gin.Context) (models.Model, error)
}

// ValidationError is validation error type
type ValidationError struct {
	Errors map[string]interface{} `json:"errors"`
}

// Validate runs validations against the given model
func Validate(v Validator, context *gin.Context) (models.Model, error) {
	model, err := v.validate(context)
	return model, err
}

// Error wraps validations error info in ValidationError object
func Error(err error) ValidationError {
	vError := ValidationError{}
	vError.Errors = make(map[string]interface{})

	for _, e := range err.(validator.ValidationErrors) {
		if e.Param != "" {
			vError.Errors[e.Field] = fmt.Sprintf("{%s: %s}", e.Tag, e.Param)
		} else {
			vError.Errors[e.Field] = fmt.Sprintf("{key: %s}", e.Tag)
		}
	}

	return vError
}
