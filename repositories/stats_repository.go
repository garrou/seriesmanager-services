package repositories

import (
	"gorm.io/gorm"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
)

type StatsRepository interface {
	FindNbSeasonsByYears(userId string) []dto.StatDto
	FindNbSeasonsByMonths(userId string) []dto.StatDto
	FindTimeSeasonsByYears(userId string) []dto.StatDto
	FindEpisodesByYears(userId string) []dto.StatDto
	FindTotalSeries(userId string) int64
	FindTotalTime(userId string) dto.SeriesStatDto
	FindTimeCurrentMonth(userId string) dto.SeriesStatDto
	FindTimeCurrentYear(userId string) dto.SeriesStatDto
	FindAddedSeriesByYears(userId string) []dto.StatDto
}

type statsRepository struct {
	db *gorm.DB
}

func NewStatsRepository(db *gorm.DB) StatsRepository {
	return &statsRepository{db: db}
}

func (s *statsRepository) FindNbSeasonsByYears(userId string) []dto.StatDto {
	var stats []dto.StatDto
	s.db.
		Model(entities.Series{}).
		Select("EXTRACT(YEAR FROM viewed_at) AS label, COUNT(*) AS value").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where("user_id = ?", userId).
		Group("label").
		Order("label").
		Scan(&stats)
	return stats
}

func (s *statsRepository) FindNbSeasonsByMonths(userId string) []dto.StatDto {
	var stats []dto.StatDto
	s.db.
		Model(entities.Series{}).
		Select("TO_CHAR(viewed_at, 'Mon') AS label, COUNT(*) AS value, EXTRACT(MONTH FROM viewed_at) AS month").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where("user_id = ?", userId).
		Group("label, month").
		Order("month").
		Scan(&stats)
	return stats
}

func (s *statsRepository) FindTimeSeasonsByYears(userId string) []dto.StatDto {
	var stats []dto.StatDto
	s.db.
		Model(entities.Series{}).
		Select("EXTRACT(YEAR FROM viewed_at) AS label, SUM(episode_length * episodes) AS value").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where("user_id = ?", userId).
		Group("label").
		Order("label").
		Scan(&stats)
	return stats
}

func (s *statsRepository) FindEpisodesByYears(userId string) []dto.StatDto {
	var stats []dto.StatDto
	s.db.
		Model(entities.Series{}).
		Select("EXTRACT(YEAR FROM viewed_at) AS label, SUM(episodes) AS value").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where("user_id = ?", userId).
		Group("label").
		Order("label").
		Scan(&stats)
	return stats
}

func (s *statsRepository) FindTotalSeries(userId string) int64 {
	var total int64
	s.db.
		Model(entities.Series{}).
		Where("user_id = ?", userId).
		Count(&total)
	return total
}

func (s *statsRepository) FindTotalTime(userId string) dto.SeriesStatDto {
	var stats dto.SeriesStatDto
	s.db.
		Model(entities.Series{}).
		Select("SUM(seasons.episodes * series.episode_length) AS total").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where("user_id = ?", userId).
		Scan(&stats)
	return stats
}

func (s *statsRepository) FindTimeCurrentMonth(userId string) dto.SeriesStatDto {
	var stats dto.SeriesStatDto
	s.db.
		Model(entities.Series{}).
		Select("SUM(seasons.episodes * series.episode_length) AS total").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where(`user_id = ? AND viewed_at >= DATE_TRUNC('month', CURRENT_DATE)`, userId).
		Scan(&stats)
	return stats
}

func (s *statsRepository) FindTimeCurrentYear(userId string) dto.SeriesStatDto {
	var stats dto.SeriesStatDto
	s.db.
		Model(entities.Series{}).
		Select("SUM(seasons.episodes * series.episode_length) AS total").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where(`user_id = ? 
				AND EXTRACT(year from viewed_at) = EXTRACT(year from now())`, userId).
		Scan(&stats)
	return stats
}

func (s *statsRepository) FindAddedSeriesByYears(userId string) []dto.StatDto {
	var stats []dto.StatDto
	s.db.
		Model(entities.Series{}).
		Select("EXTRACT(YEAR FROM added_at) AS label, COUNT(*) AS value").
		Where("user_id = ?", userId).
		Group("label").
		Scan(&stats)
	return stats
}
