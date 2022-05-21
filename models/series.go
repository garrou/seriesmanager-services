package models

// SearchSeries represents the result when search
type SearchSeries struct {
	Series []struct {
		Id     int         `json:"id"`
		Title  string      `json:"title"`
		Images interface{} `json:"images"`
	} `json:"shows"`
}

// PreviewSeries represents the result when get by id
type PreviewSeries struct {
	Series struct {
		Id     int         `json:"id"`
		Title  string      `json:"title"`
		Images interface{} `json:"images"`
	} `json:"show"`
}

// Series represents a series in database
type Series struct {
	Id   int
	User string `gorm:"column:fk_user"`
}
