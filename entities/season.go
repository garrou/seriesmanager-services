package entities

import (
	"time"
)

// Season represents a season in database
type Season struct {
	ID       int    `gorm:"autoIncrement;"`
	Number   int    `gorm:"not null;"`
	Episodes int    `gorm:"not null;"`
	Image    string `gorm:"varchar(150);"`
	ViewedAt time.Time
	SeriesID int
}
