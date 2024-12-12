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
	user, err := h.AuthService.Register(input)
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

	// Return the created user (excluding the password)
	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}
