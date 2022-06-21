package models

import (
	"time"
)

// Season represents a season in database
type Season struct {
	ID       int       `json:"id" gorm:"autoIncrement;"`
	Number   int       `json:"number" gorm:"not null;"`
	Episodes int       `json:"episodes" gorm:"not null;"`
	Image    string    `json:"image" gorm:"varchar(150);"`
	ViewedAt time.Time `json:"viewedAt"`
	SeriesID int       `json:"seriesId"`
}

// SeasonInfos represents user season infos
type SeasonInfos struct {
	ViewedAt time.Time `json:"viewedAt"`
	Duration int       `json:"duration"`
}

// SeasonStat represents number of seasons by years
type SeasonStat struct {
	Viewed int `json:"viewedAt"`
	Num    int `json:"num"`
}

// SeasonMonthStat represents number of seasons by months
type SeasonMonthStat struct {
	Month string `json:"month"`
	Num   int    `json:"num"`
}

// SeasonDetailsViewed represents number of time seasons are viewed
type SeasonDetailsViewed struct {
	Number int `json:"number"`
	Total  int `json:"total"`
}
