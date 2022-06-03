package dto

type SeasonDto struct {
	Id         int    `json:"id"`
	Number     int    `json:"number"`
	Episodes   int    `json:"episodes"`
	Image      string `json:"image"`
	StartedAt  string `json:"startedAt"`
	FinishedAt string `json:"finishedAt"`
	Series     int    `json:"seriesId"`
}

type SeasonCreateDto struct {
	Number     int    `json:"number"`
	Episodes   int    `json:"episodes"`
	Image      string `json:"image"`
	StartedAt  string `json:"startedAt"`
	FinishedAt string `json:"finishedAt"`
	Series     int    `json:"seriesId"`
}
