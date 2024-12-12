package models

import (
	"time"

	_ "gorm.io/gorm" // Add blank import to suppress unused import error
)

// User represents a user in the system.
type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"not null"`
	Email     string         `gorm:"not null;unique"`
	Password  string         `gorm:"not null"`
	RoleID    uint           `gorm:"not null"` // Foreign key to Role
	Role      Role           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	Loans     []Loan         `gorm:"foreignKey:UserID"`
}