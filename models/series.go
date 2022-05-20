package models

type SearchSeries struct {
	Series []struct {
		Id     int         `json:"id"`
		Title  string      `json:"title"`
		Images interface{} `json:"images"`
	} `json:"shows"`
}

type PreviewSeries struct {
	Series struct {
		Id     int         `json:"id"`
		Title  string      `json:"title"`
		Images interface{} `json:"images"`
	} `json:"show"`
}

type Series struct {
	Id   int
	User string `gorm:"column:fk_user"`
}
