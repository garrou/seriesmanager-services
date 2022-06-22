package services

import (
	"seriesmanager-services/models"
	"seriesmanager-services/repositories"
)

type StatsService interface {
	GetNbSeasonsByMonths(userId string) []models.SeasonMonthStat
	GetNbSeasonsByYears(userId string) []models.SeasonStat
	GetTimeSeasonsByYears(userId string) []models.SeasonStat
	GetEpisodesByYears(userId string) []models.SeasonStat
	GetTotalSeries(userId string) int64
	GetTotalTime(userId string) models.SeriesStat
	GetTimeCurrentMonth(userId string) models.SeriesStat
	GetAddedSeriesByYears(userId string) []models.SeriesAddedYears
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

func (s *statsService) GetTimeCurrentMonth(userId string) models.SeriesStat {
	return s.statsRepository.FindTimeCurrentMonth(userId)
}

func (s *statsService) GetAddedSeriesByYears(userId string) []models.SeriesAddedYears {
	return s.statsRepository.FindAddedSeriesByYears(userId)
}

func (s *statsService) GetNbSeasonsByMonths(userId string) []models.SeasonMonthStat {
	return s.statsRepository.FindNbSeasonsByMonths(userId)
}
