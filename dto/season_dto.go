package dto

import "time"

// SeasonCreateDto represents a season to create
type SeasonCreateDto struct {
	Number     int       `json:"number" binding:"required"`
	Episodes   int       `json:"episodes" binding:"required"`
	Image      string    `json:"image" validate:"max:150"`
	StartedAt  time.Time `json:"startedAt" binding:"required"`
	FinishedAt time.Time `json:"finishedAt" binding:"required"`
	SeriesId   int       `json:"sid" binding:"required"`
}

// SearchSeasonsDto represents an api season
type SearchSeasonsDto struct {
	Seasons []struct {
		Id       int    `json:"id"`
		Number   int    `json:"number"`
		Episodes int    `json:"episodes"`
		Image    string `json:"image"`
	} `json:"seasons"`
}

// SeasonsCreateAllDto represents all seasons of series to create
type SeasonsCreateAllDto struct {
	Start   time.Time `json:"start" binding:"required"`
	End     time.Time `json:"end" binding:"required"`
	Seasons []struct {
		Number   int    `json:"number" binding:"required"`
		Episodes int    `json:"episodes" binding:"required"`
		Image    string `json:"image" binding:"required"`
	} `json:"seasons" binding:"required"`
}
