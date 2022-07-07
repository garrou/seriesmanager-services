package dto

import "time"

// SeasonDto represents season
type SeasonDto struct {
	ID       int       `json:"id"`
	Number   int       `json:"number"`
	Episodes int       `json:"episodes"`
	Image    string    `json:"image"`
	ViewedAt time.Time `json:"viewedAt"`
	SeriesID int       `json:"seriesId"`
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

// SeasonsCreateDto represents all seasons of series to create
type SeasonsCreateDto struct {
	ViewedAt time.Time `json:"viewedAt" binding:"required"`
	SeriesId int       `json:"seriesId" binding:"required"`
	Seasons  []struct {
		Number   int    `json:"number" binding:"required"`
		Episodes int    `json:"episodes" binding:"required"`
		Image    string `json:"image" binding:"required"`
	} `json:"seasons" binding:"required"`
}

// SeasonInfosDto represents user season infos
type SeasonInfosDto struct {
	Id       int       `json:"id"`
	ViewedAt time.Time `json:"viewedAt"`
	Duration int       `json:"duration"`
}

// SeasonUpdateDto represent an update season
type SeasonUpdateDto struct {
	Id       int       `json:"id"`
	ViewedAt time.Time `json:"viewedAt"`
}
