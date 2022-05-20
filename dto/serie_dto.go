package dto

type SeriesCreateDto struct {
	Id   int    `json:"id" binding:"required"`
	User string `json:"omitempty"`
}
