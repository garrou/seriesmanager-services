package dto

// Episode represents an api episode
type Episode struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Season      int    `json:"season"`
	Episode     int    `json:"episode"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

type SearchEpisodes struct {
	Episodes []Episode `json:"episodes"`
}
