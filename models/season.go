package models

import "time"

// Season represents a season in database
type Season struct {
	Id         string    `json:"id"`
	Number     int       `json:"number"`
	Episodes   int       `json:"episodes"`
	Image      string    `json:"image"`
	Series     string    `json:"sid" gorm:"column:fk_series"`
	StartedAt  time.Time `json:"startedAt" gorm:"column:started_at"`
	FinishedAt time.Time `json:"finishedAt" gorm:"column:finished_at"`
}

// SeasonInfos represents user season infos
type SeasonInfos struct {
	StartedAt  time.Time `json:"startedAt"`
	FinishedAt time.Time `json:"finishedAt"`
	Duration   int       `json:"duration"`
}
