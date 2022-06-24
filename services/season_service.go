package services

import (
	"encoding/json"
	"fmt"
	"os"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
	"seriesmanager-services/helpers"
	"seriesmanager-services/repositories"
	"strconv"
)

type SeasonService interface {
	AddSeason(season dto.SeasonCreateDto) interface{}
	GetDistinctBySeriesId(seriesId string) []dto.SeasonDto
	GetInfosBySeasonBySeriesId(seriesId, number string) []dto.SeasonInfosDto
	GetDetailsSeasonsNbViewed(userId, seriesId string) []dto.StatDto
	AddAllSeasonsBySeries(userId, seriesId string, seasons dto.SeasonsCreateAllDto) interface{}
	GetToContinue(userId string) []dto.SeriesToContinueDto
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
	return s.seasonRepository.Save(entities.Season{
		Number:   season.Number,
		Episodes: season.Episodes,
		Image:    season.Image,
		ViewedAt: season.ViewedAt,
		SeriesID: season.SeriesId,
	})
}

func (s *seasonService) GetDistinctBySeriesId(seriesId string) []dto.SeasonDto {
	var seasonsDto []dto.SeasonDto
	seasons := s.seasonRepository.FindDistinctBySeriesId(seriesId)

	for _, season := range seasons {
		seasonsDto = append(seasonsDto, dto.SeasonDto{
			ID:       season.ID,
			SeriesID: season.SeriesID,
			ViewedAt: season.ViewedAt,
			Episodes: season.Episodes,
			Image:    season.Image,
			Number:   season.Number,
		})
	}
	return seasonsDto
}

func (s *seasonService) GetInfosBySeasonBySeriesId(seriesId, number string) []dto.SeasonInfosDto {
	return s.seasonRepository.FindInfosBySeriesIdBySeason(seriesId, number)
}

func (s *seasonService) GetDetailsSeasonsNbViewed(userId, seriesId string) []dto.StatDto {
	return s.seasonRepository.FindDetailsSeasonsNbViewed(userId, seriesId)
}

func (s *seasonService) AddAllSeasonsBySeries(userId, seriesId string, seasons dto.SeasonsCreateAllDto) interface{} {
	exists := s.seriesRepository.ExistsByUserIdSeriesId(userId, seriesId)
	id, err := strconv.Atoi(seriesId)

	if !exists || err != nil {
		return false
	}

	for _, season := range seasons.Seasons {
		s.seasonRepository.Save(entities.Season{
			Number:   season.Number,
			Episodes: season.Episodes,
			Image:    season.Image,
			ViewedAt: seasons.ViewedAt,
			SeriesID: id,
		})
	}
	return seasons
}

func (s *seasonService) GetToContinue(userId string) []dto.SeriesToContinueDto {
	apiKey := os.Getenv("API_KEY")
	series := s.seriesRepository.FindByUserId(userId)
	var seasons dto.SearchSeasonsDto
	var toContinue []dto.SeriesToContinueDto

	for _, userSeries := range series {
		userSeasons := s.seasonRepository.FindDistinctBySeriesId(strconv.Itoa(userSeries.ID))
		body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/seasons?id=%d&key=%s", userSeries.Sid, apiKey))

		if err := json.Unmarshal(body, &seasons); err != nil {
			panic(err.Error())
		}

		diff := len(seasons.Seasons) - len(userSeasons)

		if diff > 0 {
			toContinue = append(toContinue, dto.SeriesToContinueDto{
				Id:            userSeries.ID,
				Title:         userSeries.Title,
				Poster:        userSeries.Poster,
				EpisodeLength: userSeries.EpisodeLength,
				Sid:           userSeries.Sid,
				NbMissing:     diff,
			})
		}
	}
	return toContinue
}
