package repositories

import (
	"gorm.io/gorm"
	"seriesmanager-services/models"
)

type SeasonRepository interface {
	FindDistinctBySid(sid string) []models.Season
	Save(series models.Season) models.Season
	FindInfosBySeriesBySeason(sid, number string) []models.SeasonInfos
}

type seasonRepository struct {
	db *gorm.DB
}

func NewSeasonRepository(db *gorm.DB) SeasonRepository {
	return &seasonRepository{db: db}
}

func (s *seasonRepository) FindDistinctBySid(sid string) []models.Season {
	var seasons []models.Season
	s.db.
		Distinct("ON (number) number, *").
		Order("number, started_at").
		Find(&seasons, "fk_series = ?", sid)
	return seasons
}

func (s *seasonRepository) Save(season models.Season) models.Season {
	s.db.Save(&season)
	return season
}

func (s *seasonRepository) FindInfosBySeriesBySeason(sid, number string) []models.SeasonInfos {
	var infos []models.SeasonInfos
	s.db.
		Model(&models.Season{}).
		Select("seasons.started_at, seasons.finished_at, seasons.episodes * series.episode_length AS duration").
		Joins("JOIN series ON sid = fk_series").
		Where("fk_series = ? AND number = ?", sid, number).
		Group("started_at, finished_at, episodes, episode_length").
		Scan(&infos)
	return infos
}
