package repositories

import (
	"gorm.io/gorm"
	"seriesmanager-services/models"
)

type StatsRepository interface {
	FindNbSeasonsByYears(userId string) []models.SeasonStat
	FindTimeSeasonsByYears(userId string) []models.SeasonStat
}

type statsRepository struct {
	db *gorm.DB
}

func NewStatsRepository(db *gorm.DB) StatsRepository {
	return &statsRepository{db: db}
}

func (s *statsRepository) FindNbSeasonsByYears(userId string) []models.SeasonStat {
	var stats []models.SeasonStat
	s.db.
		Model(models.Series{}).
		Select("EXTRACT(YEAR FROM started_at) AS started, EXTRACT(YEAR FROM finished_at) AS finished, COUNT(*) AS num").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where("user_id = ?", userId).
		Group("started, finished").
		Order("started").
		Scan(&stats)
	return stats
}

func (s *statsRepository) FindTimeSeasonsByYears(userId string) []models.SeasonStat {
	var stats []models.SeasonStat
	s.db.
		Model(models.Series{}).
		Select("EXTRACT(YEAR FROM started_at) AS started, EXTRACT(YEAR FROM finished_at) AS finished, SUM(episode_length) * episodes AS num").
		Joins("JOIN seasons ON Seasons.series_id = series.id").
		Where("user_id = ?", userId).
		Group("started, finished, episodes").
		Order("started").
		Scan(&stats)
	return stats
}
