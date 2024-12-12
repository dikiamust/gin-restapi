package services

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
	"go-restapi-gin/internal/models"
)

type AuthService struct {
	DB *gorm.DB
}

// RegisterInput represents the data required for user registration
type RegisterInput struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	RoleID   uint   `json:"role_id" binding:"required"`
}

// Register creates a new user in the database
func (s *AuthService) Register(input RegisterInput) (*models.User, error) {
	// Check if email already exists
	var existingUser models.User
	if err := s.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("email already in use")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Error checking for existing user:", err)
		return nil, errors.New("internal server error")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return nil, errors.New("internal server error")
	}

	// Create the user record
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		RoleID:   input.RoleID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.DB.Create(&user).Error; err != nil {
		log.Println("Error creating user:", err)
		return nil, errors.New("internal server error")
	}

	// Mask the password before returning
	user.Password = ""
	return &user, nil
}
