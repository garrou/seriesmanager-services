package repositories

import (
	"gorm.io/gorm"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
)

type SeasonRepository interface {
	FindDistinctBySeriesId(seriesId string) []entities.Season
	Save(season entities.Season) entities.Season
	FindInfosBySeriesIdBySeason(seriesId, number string) []dto.SeasonInfosDto
	FindDetailsSeasonsNbViewed(userId, seriesId string) []dto.StatDto
}

type seasonRepository struct {
	db *gorm.DB
}

func NewSeasonRepository(db *gorm.DB) SeasonRepository {
	return &seasonRepository{db: db}
}

func (s *seasonRepository) FindDistinctBySeriesId(seriesId string) []entities.Season {
	var seasons []entities.Season
	s.db.
		Distinct("ON (number) number, *").
		Order("number, viewed_at").
		Find(&seasons, "series_id = ?", seriesId)
	return seasons
}

func (s *seasonRepository) Save(season entities.Season) entities.Season {
	s.db.Save(&season)
	return season
}

func (s *seasonRepository) FindInfosBySeriesIdBySeason(seriesId, number string) []dto.SeasonInfosDto {
	var infos []dto.SeasonInfosDto
	s.db.
		Model(&entities.Season{}).
		Select("seasons.viewed_at, seasons.episodes * series.episode_length AS duration").
		Joins("JOIN series ON series.id = series_id").
		Where("series_id = ? AND number = ?", seriesId, number).
		Group("viewed_at, episodes, episode_length").
		Scan(&infos)
	return infos
}

func (s *seasonRepository) FindDetailsSeasonsNbViewed(userId, seriesId string) []dto.StatDto {
	var details []dto.StatDto
	s.db.
		Model(entities.Season{}).
		Select("number AS label, COUNT(*) AS value").
		Joins("JOIN series ON seasons.series_id = series.id").
		Where("user_id = ? AND seasons.series_id = ?", userId, seriesId).
		Order("number").
		Group("number").
		Scan(&details)
	return details
}
