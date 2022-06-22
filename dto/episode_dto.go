package dto

// SearchEpisodesDto represents api episodes
type SearchEpisodesDto struct {
	Episodes []struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Season      int    `json:"season"`
		Episode     int    `json:"episode"`
		Code        string `json:"code"`
		Description string `json:"description"`
		Date        string `json:"date"`
	} `json:"episodes"`
}
