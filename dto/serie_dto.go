package dto

type SeriesCreateDto struct {
	Id     int    `json:"id" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Poster string `json:"poster"`
	User   string
}

type SeriesPreviewDto struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Poster string `json:"poster"`
}
