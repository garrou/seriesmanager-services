package models

import (
	"time"
)

// Season represents a season in database
type Season struct {
	ID         int       `json:"id" gorm:"autoIncrement;"`
	Number     int       `json:"number" gorm:"not null;"`
	Episodes   int       `json:"episodes" gorm:"not null;"`
	Image      string    `json:"image" gorm:"varchar(150);"`
	StartedAt  time.Time `json:"startedAt"`
	FinishedAt time.Time `json:"finishedAt"`
	SeriesID   int       `json:"seriesId"`
}

// SeasonInfos represents user season infos
type SeasonInfos struct {
	StartedAt  time.Time `json:"startedAt"`
	FinishedAt time.Time `json:"finishedAt"`
	Duration   int       `json:"duration"`
}
