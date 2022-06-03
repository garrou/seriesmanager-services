package dto

type SeasonCreateDto struct {
	Number     int    `json:"number" binding:"required"`
	Episodes   int    `json:"episodes" binding:"required"`
	Image      string `json:"image"`
	StartedAt  string `json:"startedAt" binding:"required"`
	FinishedAt string `json:"finishedAt" binding:"required"`
	Series     string `json:"sid" binding:"required"`
}
