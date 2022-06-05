package services

import (
	"seriesmanager-services/dto"
	"seriesmanager-services/models"
	"seriesmanager-services/repositories"
	"time"
)

type SeasonService interface {
	AddSeason(season dto.SeasonCreateDto) interface{}
	GetDistinctBySeriesId(seriesId string) []models.Season
	GetInfosBySeasonBySeriesId(seriesId, number string) []models.SeasonInfos
}

type seasonService struct {
	seasonRepository repositories.SeasonRepository
}

func NewSeasonService(seasonRepository repositories.SeasonRepository) SeasonService {
	return &seasonService{
		seasonRepository: seasonRepository,
	}
}

func (s *seasonService) AddSeason(season dto.SeasonCreateDto) interface{} {
	start, errStart := time.Parse(time.RFC3339, season.StartedAt)
	finish, errFinish := time.Parse(time.RFC3339, season.FinishedAt)

	if errStart != nil || errFinish != nil || start.After(finish) {
		return false
	}
	return s.seasonRepository.Save(models.Season{
		Number:     season.Number,
		Episodes:   season.Episodes,
		Image:      season.Image,
		StartedAt:  start,
		FinishedAt: finish,
		SeriesID:   season.SeriesId,
	})
}

func (s *seasonService) GetDistinctBySeriesId(seriesId string) []models.Season {
	return s.seasonRepository.FindDistinctBySeriesId(seriesId)
}

func (s *seasonService) GetInfosBySeasonBySeriesId(seriesId, number string) []models.SeasonInfos {
	return s.seasonRepository.FindInfosBySeriesIdBySeason(seriesId, number)
}
