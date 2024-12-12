package services

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"

	"go-restapi-gin/internal/models"
)

type RoleService struct {
	DB *gorm.DB
}

// RoleInput represents the data required for creating or updating a role
type RoleInput struct {
	Name        string `json:"name" binding:"required,min=3,max=50"`
	Description string `json:"description" binding:"max=255"`
}

// CreateRole creates a new role in the database
func (s *RoleService) CreateRole(input RoleInput) (*models.Role, error) {
	// Check if the role name already exists
	var existingRole models.Role
	if err := s.DB.Where("name = ?", input.Name).First(&existingRole).Error; err == nil {
		return nil, errors.New("role name already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Error checking for existing role:", err)
		return nil, errors.New("internal server error")
	}

	// Create the role record
	role := models.Role{
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.DB.Create(&role).Error; err != nil {
		log.Println("Error creating role:", err)
		return nil, errors.New("internal server error")
	}

	return &role, nil
}

// GetRoles retrieves all roles from the database
func (s *RoleService) GetRoles() ([]models.Role, error) {
	var roles []models.Role
	if err := s.DB.Find(&roles).Error; err != nil {
		log.Println("Error retrieving roles:", err)
		return nil, errors.New("internal server error")
	}
	return roles, nil
}

// UpdateRole updates an existing role in the database
func (s *RoleService) UpdateRole(id uint, input RoleInput) (*models.Role, error) {
	// Find the existing role
	var role models.Role
	if err := s.DB.First(&role, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("role not found")
		}
		log.Println("Error finding role:", err)
		return nil, errors.New("internal server error")
	}

	// Update the role fields
	role.Name = input.Name
	role.Description = input.Description
	role.UpdatedAt = time.Now()

	if err := s.DB.Save(&role).Error; err != nil {
		log.Println("Error updating role:", err)
		return nil, errors.New("internal server error")
	}

	return &role, nil
}

// DeleteRole deletes a role from the database
func (s *RoleService) DeleteRole(id uint) error {
	if err := s.DB.Delete(&models.Role{}, id).Error; err != nil {
		log.Println("Error deleting role:", err)
		return errors.New("internal server error")
	}
	return nil
}
