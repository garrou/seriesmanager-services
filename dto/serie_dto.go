package dto

type SeriesCreateDto struct {
	Id     int    `json:"id" form:"id" binding:"required"`
	FkUser string `json:"omitempty"`
}
