package repositories

import (
	"seriesmanager-services/database"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
)

func FindNbSeasonsByYears(userId string) []dto.StatDto {
	var stats []dto.StatDto
	database.Db.
		Model(entities.Series{}).
		Select("EXTRACT(YEAR FROM viewed_at) AS label, COUNT(*) AS value").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where("user_id = ?", userId).
		Group("label").
		Order("label").
		Scan(&stats)
	return stats
}

func FindNbSeasonsByMonths(userId string) []dto.StatDto {
	var stats []dto.StatDto
	database.Db.
		Model(entities.Series{}).
		Select("TO_CHAR(viewed_at, 'Month') AS label, COUNT(*) AS value").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where("user_id = ?", userId).
		Group("label").
		Scan(&stats)
	return stats
}

func FindTimeSeasonsByYears(userId string) []dto.StatDto {
	var stats []dto.StatDto
	database.Db.
		Model(entities.Series{}).
		Select("EXTRACT(YEAR FROM viewed_at) AS label, SUM(episode_length * episodes) AS value").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where("user_id = ?", userId).
		Group("label").
		Order("label").
		Scan(&stats)
	return stats
}

func FindEpisodesByYears(userId string) []dto.StatDto {
	var stats []dto.StatDto
	database.Db.
		Model(entities.Series{}).
		Select("EXTRACT(YEAR FROM viewed_at) AS label, SUM(episodes) AS value").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where("user_id = ?", userId).
		Group("label").
		Order("label").
		Scan(&stats)
	return stats
}

func FindTotalSeries(userId string) int64 {
	var total int64
	database.Db.
		Model(entities.Series{}).
		Where("user_id = ?", userId).
		Count(&total)
	return total
}

func FindTotalTime(userId string) dto.SeriesStatDto {
	var stats dto.SeriesStatDto
	database.Db.
		Model(entities.Series{}).
		Select("SUM(seasons.episodes * series.episode_length) AS total").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where("user_id = ?", userId).
		Scan(&stats)
	return stats
}

func FindTimeCurrentMonth(userId string) dto.SeriesStatDto {
	var stats dto.SeriesStatDto
	database.Db.
		Model(entities.Series{}).
		Select("SUM(seasons.episodes * series.episode_length) AS total").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where(`user_id = ? AND viewed_at >= DATE_TRUNC('month', CURRENT_DATE)`, userId).
		Scan(&stats)
	return stats
}

func FindTimeCurrentYear(userId string) dto.SeriesStatDto {
	var stats dto.SeriesStatDto
	database.Db.
		Model(entities.Series{}).
		Select("SUM(seasons.episodes * series.episode_length) AS total").
		Joins("JOIN seasons ON seasons.series_id = series.id").
		Where(`user_id = ? 
				AND EXTRACT(year from viewed_at) = EXTRACT(year from now())`, userId).
		Scan(&stats)
	return stats
}

func FindAddedSeriesByYears(userId string) []dto.StatDto {
	var stats []dto.StatDto
	database.Db.
		Model(entities.Series{}).
		Select("EXTRACT(YEAR FROM added_at) AS label, COUNT(*) AS value").
		Where("user_id = ?", userId).
		Group("label").
		Scan(&stats)
	return stats
}
