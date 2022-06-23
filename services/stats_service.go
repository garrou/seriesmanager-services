package services

import (
	"seriesmanager-services/dto"
	"seriesmanager-services/repositories"
)

type StatsService interface {
	GetNbSeasonsByMonths(userId string) []dto.StatDto
	GetNbSeasonsByYears(userId string) []dto.StatDto
	GetTimeSeasonsByYears(userId string) []dto.StatDto
	GetEpisodesByYears(userId string) []dto.StatDto
	GetTotalSeries(userId string) int64
	GetTotalTime(userId string) dto.SeriesStatDto
	GetTimeCurrentMonth(userId string) dto.SeriesStatDto
	GetAddedSeriesByYears(userId string) []dto.StatDto
}

type statsService struct {
	statsRepository repositories.StatsRepository
}

func NewStatsService(statsRepository repositories.StatsRepository) StatsService {
	return &statsService{statsRepository: statsRepository}
}

func (s *statsService) GetNbSeasonsByYears(userId string) []dto.StatDto {
	return s.statsRepository.FindNbSeasonsByYears(userId)
}

func (s *statsService) GetTimeSeasonsByYears(userId string) []dto.StatDto {
	return s.statsRepository.FindTimeSeasonsByYears(userId)
}

func (s *statsService) GetEpisodesByYears(userId string) []dto.StatDto {
	return s.statsRepository.FindEpisodesByYears(userId)
}

func (s *statsService) GetTotalSeries(userId string) int64 {
	return s.statsRepository.FindTotalSeries(userId)
}

func (s *statsService) GetTotalTime(userId string) dto.SeriesStatDto {
	return s.statsRepository.FindTotalTime(userId)
}

func (s *statsService) GetTimeCurrentMonth(userId string) dto.SeriesStatDto {
	return s.statsRepository.FindTimeCurrentMonth(userId)
}

func (s *statsService) GetAddedSeriesByYears(userId string) []dto.StatDto {
	return s.statsRepository.FindAddedSeriesByYears(userId)
}

func (s *statsService) GetNbSeasonsByMonths(userId string) []dto.StatDto {
	return s.statsRepository.FindNbSeasonsByMonths(userId)
}
