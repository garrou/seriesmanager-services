package dto

// SeriesCreateDto represents a series to create
type SeriesCreateDto struct {
	Sid           int    `json:"id" binding:"required"`
	Title         string `json:"title" binding:"required" validate:"max:150"`
	Poster        string `json:"poster" validate:"max:150"`
	EpisodeLength int    `json:"length" binding:"required"`
	UserId        string
}

// SeriesPreviewDto represents the preview of the series
type SeriesPreviewDto struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Poster        string `json:"poster"`
	EpisodeLength int    `json:"length"`
	Sid           int    `json:"sid"`
}

// PreviewSeriesDto represents the result when get by id
type PreviewSeriesDto struct {
	Series struct {
		Id     int    `json:"id"`
		Title  string `json:"title"`
		Poster struct {
			Poster string `json:"poster"`
		} `json:"images"`
	} `json:"show"`
}

// DetailsDto represent details api series
type DetailsDto struct {
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
		Notes     interface{} `json:"notes"`
		Platforms struct {
			Streaming []struct {
				Name string `json:"name"`
				Logo string `json:"logo"`
			} `json:"svods"`
		} `json:"platforms"`
	}
}

// SearchedSeries represents the results of api search series by name
type SearchedSeries struct {
	Series []struct {
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
		Notes     interface{} `json:"notes"`
		Platforms struct {
			Streaming []struct {
				Name string `json:"name"`
				Logo string `json:"logo"`
			} `json:"svods"`
		} `json:"platforms"`
	} `json:"shows"`
}

type PicturesDto struct {
	Pictures []struct {
		Url string `json:"url"`
	} `json:"pictures"`
}

type SeriesToContinueDto struct {
	Title     string `json:"title"`
	NbMissing int    `json:"nbMissing"`
}
