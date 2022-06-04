package dto

// SeriesCreateDto represents a series to create
type SeriesCreateDto struct {
	Id            int    `json:"id" binding:"required"`
	Title         string `json:"title" binding:"required"`
	Poster        string `json:"poster"`
	EpisodeLength int    `json:"length" binding:"required"`
	User          string
}

// SeriesPreviewDto represents the preview of the series
type SeriesPreviewDto struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Poster        string `json:"poster"`
	EpisodeLength int    `json:"length"`
	Sid           string `json:"sid"`
}

// SearchSeries represents a searched or discovered series
type SearchSeries struct {
	Series []struct {
		Id     int    `json:"id"`
		Title  string `json:"title"`
		Images struct {
			Poster string `json:"poster"`
		} `json:"images"`
	} `json:"shows"`
}

// PreviewSeries represents the result when get by id
type PreviewSeries struct {
	Series struct {
		Id     int    `json:"id"`
		Title  string `json:"title"`
		Poster struct {
			Poster string `json:"poster"`
		} `json:"images"`
	} `json:"show"`
}

// DetailsSeries represents the search details
type DetailsSeries struct {
	Series struct {
		Id     int    `json:"id"`
		Title  string `json:"title"`
		Images struct {
			Banner string `json:"banner"`
			Box    string `json:"box"`
			Poster string `json:"poster"`
			Show   string `json:"show"`
		} `json:"images"`
		Description string `json:"description"`
		Episodes    string `json:"episodes"`
		Seasons     []struct {
			Number   int `json:"number"`
			Episodes int `json:"episodes"`
		} `json:"seasons_details"`
		Creation  string      `json:"creation"`
		Genres    interface{} `json:"genres"`
		Length    string      `json:"length"`
		Status    string      `json:"status"`
		Platforms struct {
			Streaming []struct {
				Name string `json:"name"`
				Logo string `json:"logo"`
			} `json:"svods"`
		} `json:"platforms"`
	} `json:"show"`
}
