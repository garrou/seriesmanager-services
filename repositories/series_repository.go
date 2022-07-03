package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
)

type SeriesRepository interface {
	Exists(sid int, userId string) *gorm.DB
	Save(series entities.Series) entities.Series
	FindByUserIdSeriesId(userId string, seriesId int) interface{}
	FindByUserId(userId string) []entities.Series
	FindByUserIdAndWatching(userId string) []entities.Series
	FindByUserIdAndName(userId, name string) []entities.Series
	FindInfosBySeriesId(userId string, seriesId int) dto.SeriesInfoDto
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

func (s *seriesRepository) FindByUserIdAndWatching(userId string) []entities.Series {
	var series []entities.Series
	res := s.db.
		Order("added_at DESC").
		Find(&series, "user_id = ? AND is_watching = TRUE", userId)

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

func (s *seriesRepository) FindInfosBySeriesId(userId string, seriesId int) dto.SeriesInfoDto {
	var infos dto.SeriesInfoDto
	res := s.db.
		Model(&entities.Series{}).
		Select(`
series.id, 
episode_length * SUM(episodes) AS duration, 
COUNT(*) AS seasons, 
SUM(episodes) AS episodes,
MIN(viewed_at) AS begin, 
MAX(viewed_at) AS end,
is_watching AS watching`).
		Joins("JOIN seasons ON series.id = seasons.series_id").
		Where("series.id = ? AND user_id = ?", seriesId, userId).
		Group("series.id, episode_length").
		Scan(&infos)
	fmt.Println(res)
	return infos
}

func (s *seriesRepository) DeleteByUserBySeriesId(userId string, seriesId int) bool {
	res := s.db.
		Select("Seasons").
		Where("user_id = ? AND id = ?", userId, seriesId).
		Delete(&entities.Series{ID: seriesId})
	return res.Error == nil
}

func (s *seriesRepository) FindByUserIdSeriesId(userId string, seriesId int) interface{} {
	var series entities.Series
	s.db.Find(&series, "user_id = ? AND id = ?", userId, seriesId)
	return series
}
