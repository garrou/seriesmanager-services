package repositories

import (
	"gorm.io/gorm"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
)

type SeasonRepository interface {
	Save(season entities.Season) entities.Season
	FindDistinctBySeriesId(seriesId int) []dto.SeasonDto
	FindInfosBySeriesIdBySeason(userId string, seriesId, number int) []dto.SeasonInfosDto
	FindDetailsSeasonsNbViewed(userId string, seriesId int) []dto.StatDto
	FindById(userId string, id int) interface{}
	DeleteById(seasonId int) bool
}

type seasonRepository struct {
	db *gorm.DB
}

func NewSeasonRepository(db *gorm.DB) SeasonRepository {
	return &seasonRepository{db: db}
}

func (s *seasonRepository) Save(season entities.Season) entities.Season {
	s.db.Save(&season)
	return season
}

func (s *seasonRepository) FindDistinctBySeriesId(seriesId int) []dto.SeasonDto {
	var seasons []dto.SeasonDto
	s.db.
		Model(&entities.Season{}).
		Distinct("ON (number) number, *").
		Where("series_id = ?", seriesId).
		Order("number, viewed_at").
		Scan(&seasons)
	return seasons
}

func (s *seasonRepository) FindInfosBySeriesIdBySeason(userId string, seriesId, number int) []dto.SeasonInfosDto {
	var infos []dto.SeasonInfosDto
	s.db.
		Model(&entities.Season{}).
		Select("seasons.id, seasons.viewed_at, seasons.episodes * episode_length AS duration").
		Joins("JOIN series ON series.id = series_id").
		Where("user_id = ? AND series_id = ? AND number = ?", userId, seriesId, number).
		Group("seasons.id, viewed_at, episodes, duration").
		Scan(&infos)
	return infos
}

func (s *seasonRepository) FindDetailsSeasonsNbViewed(userId string, seriesId int) []dto.StatDto {
	var details []dto.StatDto
	s.db.
		Model(entities.Season{}).
		Select("number AS label, COUNT(*) AS value").
		Joins("JOIN series ON seasons.series_id = series.id").
		Where("user_id = ? AND seasons.series_id = ?", userId, seriesId).
		Order("number").
		Group("number").
		Scan(&details)
	return details
}

func (s *seasonRepository) FindById(userId string, id int) interface{} {
	var season entities.Season
	s.db.
		Model(entities.Season{}).
		Joins("JOIN series ON series.id = seasons.series_id").
		Find(&season, "seasons.id = ? AND user_id = ?", id, userId)
	return season
}

func (s *seasonRepository) DeleteById(seasonId int) bool {
	res := s.db.Delete(&entities.Season{ID: seasonId})
	return res.Error == nil
}
