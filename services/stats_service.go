package services

import (
	"seriesmanager-services/models"
	"seriesmanager-services/repositories"
)

type StatsService interface {
	GetNbSeasonsByYears(userId string) []models.SeasonStat
	GetTimeSeasonsByYears(userId string) []models.SeasonStat
	GetEpisodesByYears(userId string) []models.SeasonStat
	GetTotalSeries(userId string) int64
	GetTotalTime(userId string) models.SeriesStat
	GetCurrentTimeWeek(userId string) models.SeriesStat
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

func (s *statsService) GetEpisodesByYears(userId string) []models.SeasonStat {
	return s.statsRepository.FindEpisodesByYears(userId)
}

func (s *statsService) GetTotalSeries(userId string) int64 {
	return s.statsRepository.FindTotalSeries(userId)
}

func (s *statsService) GetTotalTime(userId string) models.SeriesStat {
	return s.statsRepository.FindTotalTime(userId)
}

func (s *statsService) GetCurrentTimeWeek(userId string) models.SeriesStat {
	return s.statsRepository.FindTimeCurrentWeek(userId)
}
