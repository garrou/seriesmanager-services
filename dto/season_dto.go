package dto

// SeasonCreateDto represents a season to create
type SeasonCreateDto struct {
	Number     int    `json:"number" binding:"required"`
	Episodes   int    `json:"episodes" binding:"required"`
	Image      string `json:"image" validate:"max:150"`
	StartedAt  string `json:"startedAt" binding:"required"`
	FinishedAt string `json:"finishedAt" binding:"required"`
	SeriesId   int    `json:"sid" binding:"required"`
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
