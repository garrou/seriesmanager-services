package dto

import "time"

// SeasonCreateDto represents a season to create
type SeasonCreateDto struct {
	Number   int       `json:"number" binding:"required"`
	Episodes int       `json:"episodes" binding:"required"`
	Image    string    `json:"image" validate:"max:150"`
	ViewedAt time.Time `json:"ViewedAt" binding:"required"`
	SeriesId int       `json:"sid" binding:"required"`
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
	ViewedAt time.Time `json:"viewedAt" binding:"required"`
	Seasons  []struct {
		Number   int    `json:"number" binding:"required"`
		Episodes int    `json:"episodes" binding:"required"`
		Image    string `json:"image" binding:"required"`
	} `json:"seasons" binding:"required"`
}
