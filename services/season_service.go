package services

import (
	"github.com/google/uuid"
	"seriesmanager-services/dto"
	"seriesmanager-services/models"
	"seriesmanager-services/repositories"
	"time"
)

type SeasonService interface {
	AddSeason(season dto.SeasonCreateDto) interface{}
	GetDistinctBySid(sid string) []models.Season
	GetInfosBySeasonBySeries(sid, number string) []models.SeasonInfos
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
		Id:         uuid.New().String(),
		Number:     season.Number,
		Episodes:   season.Episodes,
		Image:      season.Image,
		StartedAt:  start,
		FinishedAt: finish,
		Series:     season.Series,
	})
}

func (s *seasonService) GetDistinctBySid(sid string) []models.Season {
	return s.seasonRepository.FindDistinctBySid(sid)
}

func (s *seasonService) GetInfosBySeasonBySeries(sid, number string) []models.SeasonInfos {
	return s.seasonRepository.FindInfosBySeriesBySeason(sid, number)
}