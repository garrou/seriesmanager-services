package models

import (
	"time"
)

// User represents a user in database
type User struct {
	ID       string `gorm:"unique;type:varchar(50);not null;"`
	Username string `gorm:"type:varchar(50);not null;"`
	Email    string `gorm:"unique;type:varchar(255);not null;"`
	Password string `gorm:"not null;"`
	JoinedAt time.Time
	Banner   string `gorm:"type:varchar(150)"`
	Series   []Series
}

// TODO: add avatar
