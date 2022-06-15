package repositories

import (
	"gorm.io/gorm"
	"seriesmanager-services/models"
)

type SeasonRepository interface {
	FindDistinctBySeriesId(seriesId string) []models.Season
	Save(series models.Season) models.Season
	FindInfosBySeriesIdBySeason(seriesId, number string) []models.SeasonInfos
	FindDetailsSeasonsNbViewed(userId, seriesId string) []models.SeasonDetailsViewed
}

type seasonRepository struct {
	db *gorm.DB
}

func NewSeasonRepository(db *gorm.DB) SeasonRepository {
	return &seasonRepository{db: db}
}

func (s *seasonRepository) FindDistinctBySeriesId(seriesId string) []models.Season {
	var seasons []models.Season
	s.db.
		Distinct("ON (number) number, *").
		Order("number, started_at").
		Find(&seasons, "series_id = ?", seriesId)
	return seasons
}

func (s *seasonRepository) Save(season models.Season) models.Season {
	s.db.Save(&season)
	return season
}

func (s *seasonRepository) FindInfosBySeriesIdBySeason(seriesId, number string) []models.SeasonInfos {
	var infos []models.SeasonInfos
	s.db.
		Model(&models.Season{}).
		Select("seasons.started_at, seasons.finished_at, seasons.episodes * series.episode_length AS duration").
		Joins("JOIN series ON series.id = series_id").
		Where("series_id = ? AND number = ?", seriesId, number).
		Group("started_at, finished_at, episodes, episode_length").
		Scan(&infos)
	return infos
}

func (s *seasonRepository) FindDetailsSeasonsNbViewed(userId, seriesId string) []models.SeasonDetailsViewed {
	var details []models.SeasonDetailsViewed
	s.db.
		Model(models.Season{}).
		Select("number, COUNT(*) AS total").
		Joins("JOIN series ON seasons.series_id = series.id").
		Where("user_id = ? AND seasons.series_id = ?", userId, seriesId).
		Group("number").
		Scan(&details)
	return details
}
