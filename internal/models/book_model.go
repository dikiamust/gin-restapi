package models

import (
	"time"

	_ "gorm.io/gorm" // Add blank import to suppress unused import error
)

// Book represents a book in the library.
type Book struct {
	ID          uint           `gorm:"primaryKey"`
	Title       string         `gorm:"not null"`
	Author      string         `gorm:"not null"`
	Publisher   string         `gorm:"size:255"`
	ISBN        string         `gorm:"unique"` // International Standard Book Number
	Copies      int            `gorm:"not null;default:0"` // Total copies available, default is 0
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	Loans       []Loan         `gorm:"foreignKey:BookID"`
}
