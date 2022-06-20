package services

import (
	"seriesmanager-services/dto"
	"seriesmanager-services/models"
	"seriesmanager-services/repositories"
	"strconv"
)

type SeasonService interface {
	AddSeason(season dto.SeasonCreateDto) interface{}
	GetDistinctBySeriesId(seriesId string) []models.Season
	GetInfosBySeasonBySeriesId(seriesId, number string) []models.SeasonInfos
	GetDetailsSeasonsNbViewed(userId, seriesId string) []models.SeasonDetailsViewed
	AddAllSeasonsBySeries(userId, seriesId string, seasons dto.SeasonsCreateAllDto) interface{}
}

type seasonService struct {
	seasonRepository repositories.SeasonRepository
	seriesRepository repositories.SeriesRepository
}

func NewSeasonService(seasonRepository repositories.SeasonRepository, seriesRepository repositories.SeriesRepository) SeasonService {
	return &seasonService{
		seasonRepository: seasonRepository,
		seriesRepository: seriesRepository,
	}
}

func (s *seasonService) AddSeason(season dto.SeasonCreateDto) interface{} {
	if season.StartedAt.After(season.FinishedAt) {
		return false
	}
	return s.seasonRepository.Save(models.Season{
		Number:     season.Number,
		Episodes:   season.Episodes,
		Image:      season.Image,
		StartedAt:  season.StartedAt,
		FinishedAt: season.FinishedAt,
		SeriesID:   season.SeriesId,
	})
}

func (s *seasonService) GetDistinctBySeriesId(seriesId string) []models.Season {
	return s.seasonRepository.FindDistinctBySeriesId(seriesId)
}

func (s *seasonService) GetInfosBySeasonBySeriesId(seriesId, number string) []models.SeasonInfos {
	return s.seasonRepository.FindInfosBySeriesIdBySeason(seriesId, number)
}

func (s *seasonService) GetDetailsSeasonsNbViewed(userId, seriesId string) []models.SeasonDetailsViewed {
	return s.seasonRepository.FindDetailsSeasonsNbViewed(userId, seriesId)
}

func (s *seasonService) AddAllSeasonsBySeries(userId, seriesId string, seasons dto.SeasonsCreateAllDto) interface{} {
	exists := s.seriesRepository.ExistsByUserIdSeriesId(userId, seriesId)
	id, err := strconv.Atoi(seriesId)

	if !exists || seasons.Start.After(seasons.End) || err != nil {
		return false
	}
	days := int(seasons.End.Sub(seasons.Start).Hours() / 24)
	nbSeasons := len(seasons.Seasons)
	start := seasons.Start

	for _, season := range seasons.Seasons {
		daysToAdd := days / nbSeasons

		if daysToAdd < 1 {
			daysToAdd = 1
		}
		end := start.AddDate(0, 0, daysToAdd)

		s.seasonRepository.Save(models.Season{
			Number:     season.Number,
			Episodes:   season.Episodes,
			Image:      season.Image,
			StartedAt:  start,
			FinishedAt: end,
			SeriesID:   id,
		})
		start = end
	}
	return seasons
}
