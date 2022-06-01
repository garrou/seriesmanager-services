package models

type PreviewSeasons struct {
	Seasons []struct {
		Number   int    `json:"number"`
		Episodes int    `json:"episodes"`
		Image    string `json:"image"`
		FkSeries int    `json:"seriesId"`
	}
}
