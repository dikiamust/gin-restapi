package models

import (
	"time"

	_ "gorm.io/gorm" // Add blank import to suppress unused import error
)

// Role represents the role of a user in the system.
type Role struct {
	ID          uint           `gorm:"primaryKey"`
	Name        string         `gorm:"not null;unique"` // e.g., "admin", "student"
	Description string         `gorm:"size:255"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	Users       []User         `gorm:"foreignKey:RoleID"`
}
