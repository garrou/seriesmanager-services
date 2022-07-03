package entities

import (
	"time"
)

// Series represents a series in database
type Series struct {
	ID            int       `gorm:"autoIncrement;"`
	Sid           int       `gorm:"not null;"`
	Title         string    `gorm:"type:varchar(150);not null;"`
	Poster        string    `gorm:"type:varchar(150);"`
	EpisodeLength int       `gorm:"not null;"`
	AddedAt       time.Time `gorm:"not null;"`
	Seasons       []Season
	UserID        string `gorm:"not null;"`
	IsWatching    bool   `gorm:"type:boolean;default:true"`
}
