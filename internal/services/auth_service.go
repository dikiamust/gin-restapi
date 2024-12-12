package services

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
	"go-restapi-gin/internal/models"
)

type AuthService struct {
	DB           *gorm.DB
	JWTSecretKey string
}

// RegisterInput represents the data required for user registration
type RegisterInput struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	RoleID   uint   `json:"role_id" binding:"required"`
}

// LoginInput represents the data required for user login
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// TokenResponse represents the response containing the JWT token
type TokenResponse struct {
	Token string `json:"token"`
	User  *models.User `json:"user"`
}

// generateJWT generates a new JWT token for the given user
func (s *AuthService) generateJWT(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.RoleID,
		"exp":   time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
	})

	tokenString, err := token.SignedString([]byte(s.JWTSecretKey))
	if err != nil {
		log.Println("Error generating token:", err)
		return "", errors.New("internal server error")
	}

	return tokenString, nil
}

// Register creates a new user in the database and returns a JWT token
func (s *AuthService) Register(input RegisterInput) (*TokenResponse, error) {
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

	// Generate JWT token
	token, err := s.generateJWT(&user)
	if err != nil {
		return nil, err
	}

	// Mask the password before returning
	user.Password = ""
	return &TokenResponse{
		Token: token,
		User:  &user,
	}, nil
}

// Login authenticates a user and returns a JWT token if successful
func (s *AuthService) Login(input LoginInput) (*TokenResponse, error) {
	var user models.User

	// Find user by email
	if err := s.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		log.Println("Error finding user:", err)
		return nil, errors.New("internal server error")
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := s.generateJWT(&user)
	if err != nil {
		return nil, err
	}

	// Mask the password before returning
	user.Password = ""
	return &TokenResponse{
		Token: token,
		User:  &user,
	}, nil
}
