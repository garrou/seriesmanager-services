package repositories

import (
	"gorm.io/gorm"
	"seriesmanager-services/database"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
)

func SaveSeries(series entities.Series) entities.Series {
	database.Db.Save(&series)
	return series
}

func FindByUserId(userId string) []entities.Series {
	var series []entities.Series
	res := database.Db.
		Order("added_at DESC").
		Find(&series, "user_id = ?", userId)

	if res.Error == nil {
		return series
	}
	return nil
}

func FindByUserIdAndWatching(userId string) []entities.Series {
	var series []entities.Series
	res := database.Db.
		Order("title").
		Find(&series, "user_id = ? AND is_watching = TRUE", userId)

	if res.Error == nil {
		return series
	}
	return nil
}

func FindByUserIdAndName(userId, title string) []entities.Series {
	var series []entities.Series
	res := database.Db.Find(&series, "user_id = ? AND UPPER(title) LIKE UPPER(?)", userId, "%"+title+"%")

	if res.Error == nil {
		return series
	}
	return nil
}

func SeriesExists(sid int, userId string) *gorm.DB {
	var series entities.Series
	return database.Db.Take(&series, "sid = ? AND user_id = ?", sid, userId)
}

func FindInfosBySeriesId(userId string, seriesId int) dto.SeriesInfoDto {
	var infos dto.SeriesInfoDto
	database.Db.
		Model(&entities.Series{}).
		Select(`
series.id, 
episode_length * SUM(episodes) AS duration, 
COUNT(*) AS seasons, 
SUM(episodes) AS episodes,
MIN(viewed_at) AS begin, 
MAX(viewed_at) AS end,
is_watching AS watching`).
		Joins("JOIN seasons ON series.id = seasons.series_id").
		Where("series.id = ? AND user_id = ?", seriesId, userId).
		Group("series.id, episode_length").
		Scan(&infos)
	return infos
}

func DeleteByUserBySeriesId(userId string, seriesId int) bool {
	res := database.Db.
		Select("Seasons").
		Where("user_id = ? AND id = ?", userId, seriesId).
		Delete(&entities.Series{ID: seriesId})
	return res.Error == nil
}

func FindByUserIdSeriesId(userId string, seriesId int) interface{} {
	var series entities.Series
	database.Db.Find(&series, "user_id = ? AND id = ?", userId, seriesId)
	return series
}
