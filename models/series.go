package models

import (
	"time"
)

// Series represents a series in database
type Series struct {
	ID            int    `gorm:"autoIncrement;"`
	Sid           int    `gorm:"not null;"`
	Title         string `gorm:"type:varchar(150);not null;"`
	Poster        string `gorm:"type:varchar(150);"`
	EpisodeLength int    `gorm:"not null;"`
	AddedAt       time.Time
	Seasons       []Season
	UserID        string
}

// SeriesInfo represents user series info
type SeriesInfo struct {
	Duration   int       `json:"duration"`
	Seasons    int       `json:"seasons"`
	Episodes   int       `json:"episodes"`
	StartedAt  time.Time `json:"startedAt"`
	FinishedAt time.Time `json:"finishedAt"`
}
