package models

type PreviewEpisode struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Number      int         `json:"episode"`
	Season      int         `json:"season"`
	Code        string      `json:"code"`
	Description string      `json:"description"`
	Director    string      `json:"director"`
	Note        interface{} `json:"note"`
	Date        string      `json:"date"`
	Series      interface{} `json:"show"`
}

type Episode struct {
	Id   int
	User string `gorm:"column:fk_user"`
}
