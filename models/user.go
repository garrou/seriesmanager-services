package models

import (
	"time"
)

// User represents a user in database
type User struct {
	ID       string    `gorm:"unique;type:varchar(50);not null;"`
	Username string    `gorm:"type:varchar(50);not null;"`
	Email    string    `gorm:"unique;type:varchar(255);not null;"`
	Password string    `gorm:"not null;type:varchar(255);not null"`
	JoinedAt time.Time `gorm:"not null;"`
	Banner   string    `gorm:"type:varchar(150)"`
	Series   []Series
}
