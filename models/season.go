package models

import "time"

// Season represents a database season
type Season struct {
	Id         string    `json:"id"`
	Number     int       `json:"number"`
	Episodes   int       `json:"episodes"`
	Image      string    `json:"image"`
	Series     string    `json:"sid" gorm:"column:fk_series"`
	StartedAt  time.Time `json:"startedAt" gorm:"column:started_at"`
	FinishedAt time.Time `json:"finishedAt" gorm:"column:finished_at"`
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
