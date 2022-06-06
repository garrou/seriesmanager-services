package repositories

import (
	"gorm.io/gorm"
	"seriesmanager-services/models"
)

type SeriesRepository interface {
	Save(series models.Series) models.Series
	FindByUserId(userId string) []models.Series
	FindByUserIdAndTitle(userId, title string) []models.Series
	Exists(seriesId int, userId string) *gorm.DB
	FindInfosBySeriesId(seriesId int) models.SeriesInfo
	DeleteByUserBySeriesId(userId string, seriesId int) bool
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
	res := s.db.
		Order("added_at DESC").
		Find(&series, "user_id = ?", userId)

	if res.Error == nil {
		return series
	}
	return nil
}

func (s *seriesRepository) FindByUserIdAndTitle(userId, title string) []models.Series {
	var series []models.Series
	res := s.db.Find(&series, "user_id = ? AND UPPER(title) LIKE UPPER(?)", userId, "%"+title+"%")

	if res.Error == nil {
		return series
	}
	return nil
}

func (s *seriesRepository) Exists(seriesId int, userId string) *gorm.DB {
	var series models.Series
	return s.db.Take(&series, "sid = ? AND user_id = ?", seriesId, userId)
}

func (s *seriesRepository) FindInfosBySeriesId(seriesId int) models.SeriesInfo {
	var infos models.SeriesInfo
	s.db.
		Model(&models.Series{}).
		Select(`episode_length * SUM(episodes) AS duration, 
COUNT(*) AS seasons, 
SUM(episodes) AS episodes, 
MIN(started_at) AS started_at, 
MAX(finished_at) AS finished_at`).
		Joins("JOIN seasons ON series.id = seasons.series_id").
		Where("series.id = ?", seriesId).
		Group("episode_length").
		Scan(&infos)
	return infos
}

func (s *seriesRepository) DeleteByUserBySeriesId(userId string, seriesId int) bool {
	res := s.db.Select("Seasons").
		Where("user_id = ? AND id = ?", userId, seriesId).
		Delete(&models.Series{ID: seriesId})
	return res.Error == nil
}
