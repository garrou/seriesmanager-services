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
	FindInfosBySid(sid string) models.SeriesInfo
	DeleteByUserBySid(userId, sid string) bool
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
		Find(&series, "fk_user = ?", userId)

	if res.Error == nil {
		return series
	}
	return nil
}

func (s *seriesRepository) FindByUserIdAndTitle(userId, title string) []models.Series {
	var series []models.Series
	res := s.db.Find(&series, "fk_user = ? AND UPPER(title) LIKE UPPER(?)", userId, "%"+title+"%")

	if res.Error == nil {
		return series
	}
	return nil
}

func (s *seriesRepository) Exists(seriesId int, userId string) *gorm.DB {
	var series models.Series
	return s.db.Take(&series, "id = ? AND fk_user = ?", seriesId, userId)
}

func (s *seriesRepository) FindInfosBySid(sid string) models.SeriesInfo {
	var infos models.SeriesInfo
	s.db.
		Model(&models.Series{}).
		Select(`episode_length * SUM(episodes) AS duration, 
COUNT(*) AS seasons, 
SUM(episodes) AS episodes, 
MIN(started_at) AS started_at, 
MAX(finished_at) AS finished_at`).
		Joins("JOIN seasons ON sid = fk_series").
		Where("sid = ?", sid).
		Group("episode_length").
		Scan(&infos)
	return infos
}

func (s *seriesRepository) DeleteByUserBySid(userId, sid string) bool {
	res := s.db.Delete(&models.Series{}, "fk_user = ? AND sid = ?", userId, sid)
	return res.Error == nil
}
