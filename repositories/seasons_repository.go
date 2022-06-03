package repositories

import (
	"gorm.io/gorm"
	"seriesmanager-services/models"
)

type SeasonRepository interface {
	FindBySeriesId(sid int) []models.Season
	Save(series models.Season) models.Season
}

type seasonRepository struct {
	db *gorm.DB
}

func NewSeasonRepository(db *gorm.DB) SeasonRepository {
	return &seasonRepository{db: db}
}

func (s *seasonRepository) FindBySeriesId(sid int) []models.Season {
	var seasons []models.Season
	res := s.db.Where("fk_series = ?", sid).Order("number").Find(&seasons)

	if res.Error == nil {
		return seasons
	}
	return nil
}

func (s *seasonRepository) Save(season models.Season) models.Season {
	s.db.Save(&season)
	return season
}
