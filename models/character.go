package models

type Characters struct {
	Characters []struct {
		Id      string `json:"person_id"`
		Name    string `json:"name"`
		Actor   string `json:"actor"`
		Picture string `json:"picture"`
	} `json:"characters"`
}
