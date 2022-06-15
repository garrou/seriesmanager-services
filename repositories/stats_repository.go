package repositories

import (
	"gorm.io/gorm"
	"seriesmanager-services/models"
)

type StatsRepository interface {
	FindNbSeasonsByYears(userId string) []models.SeasonStat
	FindTimeSeasonsByYears(userId string) []models.SeasonStat
	FindEpisodesByYears(userId string) []models.SeasonStat
	FindTotalSeries(userId string) int64
	FindTotalTime(userId string) models.SeriesStat
	FindTimeCurrentWeek(userId string) models.SeriesStat
	FindTimeCurrentYear(userId string) models.SeriesStat
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
		Select(`EXTRACT(YEAR FROM started_at) AS started,
			EXTRACT(YEAR FROM finished_at) AS finished,
			COUNT(*) AS num`).
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
		Select(`EXTRACT(YEAR FROM started_at) AS started,
			EXTRACT(YEAR FROM finished_at) AS finished, 
			SUM(episode_length * episodes) AS num`).
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where("user_id = ?", userId).
		Group("started, finished").
		Order("started").
		Scan(&stats)
	return stats
}

func (s *statsRepository) FindEpisodesByYears(userId string) []models.SeasonStat {
	var stats []models.SeasonStat
	s.db.
		Model(models.Series{}).
		Select(`EXTRACT(YEAR FROM started_at) AS started,
			EXTRACT(YEAR FROM finished_at) AS finished, 
			SUM(episodes) AS num`).
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where("user_id = ?", userId).
		Group("started, finished").
		Order("started").
		Scan(&stats)
	return stats
}

func (s *statsRepository) FindTotalSeries(userId string) int64 {
	var total int64
	s.db.
		Model(models.Series{}).
		Count(&total).
		Where("user_id = ?", userId)
	return total
}

func (s *statsRepository) FindTotalTime(userId string) models.SeriesStat {
	var stats models.SeriesStat
	s.db.
		Model(models.Series{}).
		Select("SUM(seasons.episodes * series.episode_length) AS total").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where("user_id = ?", userId).
		Scan(&stats)
	return stats
}

func (s *statsRepository) FindTimeCurrentWeek(userId string) models.SeriesStat {
	var stats models.SeriesStat
	s.db.
		Model(models.Series{}).
		Select("SUM(seasons.episodes * series.episode_length) AS total").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where("user_id = ? AND started_at >= NOW()::DATE - 7", userId).
		Scan(&stats)
	return stats
}

func (s *statsRepository) FindTimeCurrentYear(userId string) models.SeriesStat {
	var stats models.SeriesStat
	s.db.
		Model(models.Series{}).
		Select("SUM(seasons.episodes * series.episode_length) AS total").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where(`user_id = ? 
				AND EXTRACT(year from started_at) = EXTRACT(year from now()) 
				AND EXTRACT(year from finished_at) = EXTRACT(year from now())`, userId).
		Scan(&stats)
	return stats
}
