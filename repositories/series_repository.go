package repositories

import (
	"gorm.io/gorm"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
)

type SeriesRepository interface {
	Save(series entities.Series) entities.Series
	ExistsByUserIdSeriesId(userId, seriesId string) bool
	FindByUserId(userId string) []entities.Series
	FindByUserIdAndName(userId, name string) []entities.Series
	Exists(sid int, userId string) *gorm.DB
	FindInfosBySeriesId(seriesId string) dto.SeriesInfoDto
	DeleteByUserBySeriesId(userId string, seriesId int) bool
}

type seriesRepository struct {
	db *gorm.DB
}

func NewSeriesRepository(db *gorm.DB) SeriesRepository {
	return &seriesRepository{db: db}
}

func (s *seriesRepository) Save(series entities.Series) entities.Series {
	s.db.Save(&series)
	return series
}

func (s *seriesRepository) FindByUserId(userId string) []entities.Series {
	var series []entities.Series
	res := s.db.
		Order("added_at DESC").
		Find(&series, "user_id = ?", userId)

	if res.Error == nil {
		return series
	}
	return nil
}

func (s *seriesRepository) FindByUserIdAndName(userId, title string) []entities.Series {
	var series []entities.Series
	res := s.db.Find(&series, "user_id = ? AND UPPER(title) LIKE UPPER(?)", userId, "%"+title+"%")

	if res.Error == nil {
		return series
	}
	return nil
}

func (s *seriesRepository) Exists(sid int, userId string) *gorm.DB {
	var series entities.Series
	return s.db.Take(&series, "sid = ? AND user_id = ?", sid, userId)
}

func (s *seriesRepository) FindInfosBySeriesId(seriesId string) dto.SeriesInfoDto {
	var infos dto.SeriesInfoDto
	s.db.
		Model(&entities.Series{}).
		Select(`episode_length * SUM(episodes) AS duration, 
COUNT(*) AS seasons, 
SUM(episodes) AS episodes,
MIN(viewed_at) AS begin, 
MAX(viewed_at) AS end`).
		Joins("JOIN seasons ON series.id = seasons.series_id").
		Where("series.id = ?", seriesId).
		Group("episode_length").
		Scan(&infos)
	return infos
}

func (s *seriesRepository) DeleteByUserBySeriesId(userId string, seriesId int) bool {
	res := s.db.
		Select("Seasons").
		Where("user_id = ? AND id = ?", userId, seriesId).
		Delete(&entities.Series{ID: seriesId})
	return res.Error == nil
}

func (s *seriesRepository) ExistsByUserIdSeriesId(userId, seriesId string) bool {
	var series entities.Series
	res := s.db.Take(&series, "user_id = ? AND id = ?", userId, seriesId)
	return res.Error == nil
}
