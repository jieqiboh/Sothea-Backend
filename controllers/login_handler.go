package controllers

import (
	"github.com/gin-gonic/gin"
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
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
