package models

import (
	"time"

	_ "gorm.io/gorm" // Add blank import to suppress unused import error
)

// Loan represents the record of a book loan.
type Loan struct {
	ID        uint           `gorm:"primaryKey"`
	UserID    uint           `gorm:"not null"` // Foreign key to User
	User      User           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BookID    uint           `gorm:"not null"` // Foreign key to Book
	Book      Book           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LoanDate  time.Time      `gorm:"not null"`
	ReturnDue time.Time      `gorm:"not null"` // Due date for returning the book
	Returned  bool           `gorm:"default:false"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
}

