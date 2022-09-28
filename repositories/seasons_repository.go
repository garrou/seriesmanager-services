package repositories

import (
	"seriesmanager-services/database"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
)

func SaveSeason(season entities.Season) entities.Season {
	database.Db.Save(&season)
	return season
}

func FindDistinctBySeriesId(seriesId int) []entities.Season {
	var seasons []entities.Season
	database.Db.
		Distinct("ON (number) number, *").
		Order("number, viewed_at").
		Find(&seasons, "series_id = ?", seriesId)
	return seasons
}

func FindInfosBySeriesIdBySeason(userId string, seriesId, number int) []dto.SeasonInfosDto {
	var infos []dto.SeasonInfosDto
	database.Db.
		Model(&entities.Season{}).
		Select("seasons.id, seasons.viewed_at, seasons.episodes * episode_length AS duration").
		Joins("JOIN series ON series.id = series_id").
		Where("user_id = ? AND series_id = ? AND number = ?", userId, seriesId, number).
		Group("seasons.id, viewed_at, episodes, duration").
		Scan(&infos)
	return infos
}

func FindDetailsSeasonsNbViewed(userId string, seriesId int) []dto.StatDto {
	var details []dto.StatDto
	database.Db.
		Model(entities.Season{}).
		Select("number AS label, COUNT(*) AS value").
		Joins("JOIN series ON seasons.series_id = series.id").
		Where("user_id = ? AND seasons.series_id = ?", userId, seriesId).
		Order("number").
		Group("number").
		Scan(&details)
	return details
}

func FindSeasonById(userId string, id int) interface{} {
	var season entities.Season
	database.Db.
		Joins("JOIN series ON series.id = seasons.series_id").
		Find(&season, "seasons.id = ? AND user_id = ?", id, userId)
	return season
}

func DeleteSeasonById(seasonId int) bool {
	res := database.Db.Delete(&entities.Season{ID: seasonId})
	return res.Error == nil
}
