package models

import "time"

// Series represents a series in database
type Series struct {
	Id            int
	Title         string
	Poster        string
	EpisodeLength int       `gorm:"column:episode_length"`
	AddedAt       time.Time `gorm:"column:added_at"`
	User          string    `gorm:"column:fk_user"`
	Sid           string
}

// SeriesInfo represents user series info
type SeriesInfo struct {
	Duration   int       `json:"duration"`
	Seasons    int       `json:"seasons"`
	Episodes   int       `json:"episodes"`
	StartedAt  time.Time `json:"startedAt"`
	FinishedAt time.Time `json:"finishedAt"`
}
