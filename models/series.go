package models

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

// Series represents a series in database
type Series struct {
	Id            int
	Title         string
	Poster        string
	EpisodeLength int    `gorm:"column:episode_length"`
	User          string `gorm:"column:fk_user"`
	Sid           int    `gorm:"autoIncrement"`
}
