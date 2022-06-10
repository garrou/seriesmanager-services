package services

import (
	"seriesmanager-services/models"
	"seriesmanager-services/repositories"
)

type StatsService interface {
	GetNbSeasonsByYears(userId string) []models.SeasonStat
	GetTimeSeasonsByYears(userId string) []models.SeasonStat
}

type statsService struct {
	statsRepository repositories.StatsRepository
}

func NewStatsService(statsRepository repositories.StatsRepository) StatsService {
	return &statsService{statsRepository: statsRepository}
}

func (s *statsService) GetNbSeasonsByYears(userId string) []models.SeasonStat {
	return s.statsRepository.FindNbSeasonsByYears(userId)
}

func (s *statsService) GetTimeSeasonsByYears(userId string) []models.SeasonStat {
	return s.statsRepository.FindTimeSeasonsByYears(userId)
}
