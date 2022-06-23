package models

type Pictures struct {
	Pictures []struct {
		Url string `json:"url"`
	} `json:"pictures"`
}
