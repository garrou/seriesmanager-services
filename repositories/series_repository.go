package repositories

import (
	"gorm.io/gorm"
	"services-series-manager/models"
)

type SeriesRepository interface {
	Save(series models.Series) models.Series
	FindByUserId(userId string) []models.Series
	Exists(seriesId int, userId string) *gorm.DB
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

func (s *seriesRepository) Exists(seriesId int, userId string) *gorm.DB {
	var series models.Series
	return s.db.Where("id = ? and fk_user = ?", seriesId, userId).Take(&series)
}
