package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jieqiboh/sothea_backend/entities"
	"net/http"
)

// LoginHandler represent the httphandler for patient
type LoginHandler struct {
	Usecase entities.LoginUseCase
}

// NewLoginHandler will initialize the resources endpoint
func NewLoginHandler(e *gin.Engine, us entities.LoginUseCase) {
	handler := &LoginHandler{
		Usecase: us,
	}
	e.POST("/login", handler.Login)
}

func (l *LoginHandler) Login(c *gin.Context) {
	// username and password are in the json body
	var u entities.User
	if err := c.ShouldBindJSON(&u); err != nil {
		// Use type assertion to check if err is of type validator.ValidationErrors
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			// Get the first Validation Error
			fieldErr := validationErrs[0]
			c.JSON(http.StatusBadRequest, gin.H{"error": fieldErr.Error()})
			return // exit on first error
		} else {
			// Handle other types of errors (e.g., JSON binding errors)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	ctx := c.Request.Context()

	tokenString, err := l.Usecase.Login(ctx, u)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": entities.ErrLoginFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
	return
}
