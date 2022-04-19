package repositories

import (
	"gorm.io/gorm"
	"services-series-manager/models"
)

type SeriesRepository interface {
	Save(series models.Series) models.Series
	FindByUserId(userId string) []models.Series
}

type seriesRepository struct {
	db *gorm.DB
}

func NewSeriesRepository(db *gorm.DB) SeriesRepository {
	return &seriesRepository{db: db}
}

func (s *seriesRepository) Save(series models.Series) models.Series {
	s.db.Save(&series)
	return series
}

func (s *seriesRepository) FindByUserId(userId string) []models.Series {
	var series []models.Series
	res := s.db.Where("fk_user = ?", userId).Find(&series)

	if res.Error == nil {
		return series
	}
	return nil
}
