package repositories

import (
	"gorm.io/gorm"
	"seriesmanager-services/models"
)

type SeasonRepository interface {
	FindDistinctBySeriesId(seriesId string) []models.Season
	Save(series models.Season) models.Season
	FindInfosBySeriesIdBySeason(seriesId, number string) []models.SeasonInfos
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

func (s *seasonRepository) FindInfosBySeriesIdBySeason(sid, number string) []models.SeasonInfos {
	var infos []models.SeasonInfos
	s.db.
		Model(&models.Season{}).
		Select("seasons.started_at, seasons.finished_at, seasons.episodes * series.episode_length AS duration").
		Joins("JOIN series ON sid = series_id").
		Where("series_id = ? AND number = ?", sid, number).
		Group("started_at, finished_at, episodes, episode_length").
		Scan(&infos)
	return infos
}
