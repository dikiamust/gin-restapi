package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-restapi-gin/internal/services"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

// RegisterHandler handles user registration requests
func (h *AuthHandler) RegisterHandler(c *gin.Context) {
	var input services.RegisterInput

	// Bind and validate the input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input: " + err.Error(),
		})
		return
	}

	// Call the Register service
	tokenResponse, err := h.AuthService.Register(input)
	if err != nil {
		if err.Error() == "email already in use" {
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register user",
		})
		return
	}

	// Return the created user with token
	c.JSON(http.StatusCreated, tokenResponse)
}

// LoginHandler handles user login requests
func (h *AuthHandler) LoginHandler(c *gin.Context) {
	var input services.LoginInput

	// Bind and validate the input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input: " + err.Error(),
		})
		return
	}

	// Call the Login service
	tokenResponse, err := h.AuthService.Login(input)
	if err != nil {
		if err.Error() == "invalid email or password" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to login",
		})
		return
	}

	// Return the token and user information
	c.JSON(http.StatusOK, tokenResponse)
}
