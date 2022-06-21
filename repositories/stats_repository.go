package repositories

import (
	"gorm.io/gorm"
	"seriesmanager-services/models"
)

type StatsRepository interface {
	FindNbSeasonsByYears(userId string) []models.SeasonStat
	FindNbSeasonsByMonths(userId string) []models.SeasonMonthStat
	FindTimeSeasonsByYears(userId string) []models.SeasonStat
	FindEpisodesByYears(userId string) []models.SeasonStat
	FindTotalSeries(userId string) int64
	FindTotalTime(userId string) models.SeriesStat
	FindTimeCurrentMonth(userId string) models.SeriesStat
	FindTimeCurrentYear(userId string) models.SeriesStat
	FindAddedSeriesByYears(userId string) []models.SeriesAddedYears
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
		Select("EXTRACT(YEAR FROM viewed_at) AS viewed, COUNT(*) AS num").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where("user_id = ?", userId).
		Group("viewed").
		Order("viewed").
		Scan(&stats)
	return stats
}

func (s *statsRepository) FindNbSeasonsByMonths(userId string) []models.SeasonMonthStat {
	var stats []models.SeasonMonthStat
	s.db.
		Model(models.Series{}).
		Select("TO_CHAR(viewed_at, 'Month') AS month, COUNT(*) AS num").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where("user_id = ?", userId).
		Group("month").
		Scan(&stats)
	return stats
}

func (s *statsRepository) FindTimeSeasonsByYears(userId string) []models.SeasonStat {
	var stats []models.SeasonStat
	s.db.
		Model(models.Series{}).
		Select("EXTRACT(YEAR FROM viewed_at) AS viewed, SUM(episode_length * episodes) AS num").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where("user_id = ?", userId).
		Group("viewed").
		Order("viewed").
		Scan(&stats)
	return stats
}

func (s *statsRepository) FindEpisodesByYears(userId string) []models.SeasonStat {
	var stats []models.SeasonStat
	s.db.
		Model(models.Series{}).
		Select("EXTRACT(YEAR FROM viewed_at) AS viewed, SUM(episodes) AS num").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where("user_id = ?", userId).
		Group("viewed").
		Order("viewed").
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

func (s *statsRepository) FindTimeCurrentMonth(userId string) models.SeriesStat {
	var stats models.SeriesStat
	s.db.
		Model(models.Series{}).
		Select("SUM(seasons.episodes * series.episode_length) AS total").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where(`user_id = ? 
AND EXTRACT(YEAR FROM viewed_at) = EXTRACT(YEAR FROM NOW())
AND EXTRACT(MONTH FROM viewed_at) = EXTRACT(MONTH FROM NOW())`, userId).
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
				AND EXTRACT(year from viewed_at) = EXTRACT(year from now())`, userId).
		Scan(&stats)
	return stats
}

func (s *statsRepository) FindAddedSeriesByYears(userId string) []models.SeriesAddedYears {
	var stats []models.SeriesAddedYears
	s.db.
		Model(models.Series{}).
		Select("EXTRACT(YEAR FROM added_at) AS added, COUNT(*) AS total").
		Where("user_id = ?", userId).
		Group("added").
		Scan(&stats)
	return stats
}
