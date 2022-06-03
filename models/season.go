package models

import "time"

// Season represents a viewed season in database
type Season struct {
	Id         int `gorm:"autoIncrement"`
	Number     int
	Episodes   int
	Image      string
	Series     int       `gorm:"column:fk_series"`
	StartedAt  time.Time `gorm:"column:started_at"`
	FinishedAt time.Time `gorm:"column:finished_at"`
}

// SearchSeasons represents an api season
type SearchSeasons struct {
	Seasons []struct {
		Id       int    `json:"id"`
		Number   int    `json:"number"`
		Episodes int    `json:"episodes"`
		Image    string `json:"image"`
	} `json:"seasons"`
}
