package dto

type StatDto struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

// SeriesStatDto represents total duration
type SeriesStatDto struct {
	Total int `json:"total"`
}
