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
	return s.seasonRepository.Save(models.Season{
		Number:   season.Number,
		Episodes: season.Episodes,
		Image:    season.Image,
		ViewedAt: season.ViewedAt,
		SeriesID: season.SeriesId,
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

	if !exists || err != nil {
		return false
	}

	for _, season := range seasons.Seasons {
		s.seasonRepository.Save(models.Season{
			Number:   season.Number,
			Episodes: season.Episodes,
			Image:    season.Image,
			ViewedAt: seasons.ViewedAt,
			SeriesID: id,
		})
	}
	return seasons
}
