package services

import (
	"encoding/json"
	"fmt"
	"os"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
	"seriesmanager-services/helpers"
	"seriesmanager-services/repositories"
)

type SeasonService interface {
	GetDistinctBySeriesId(userId string, seriesId int) []dto.SeasonDto
	GetInfosBySeasonBySeriesId(userId string, seriesId, number int) []dto.SeasonInfosDto
	GetDetailsSeasonsNbViewed(userId string, seriesId int) []dto.StatDto
	AddSeasonsBySeries(userId string, seriesId int, seasons dto.SeasonsCreateDto) interface{}
	GetToContinue(userId string) []dto.SeriesToContinueDto
	UpdateSeason(userId string, updateDto dto.SeasonUpdateDto) interface{}
	DeleteSeason(userId string, seasonId int) bool
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

func (s *seasonService) GetDistinctBySeriesId(userId string, seriesId int) []dto.SeasonDto {
	res := s.seriesRepository.FindByUserIdSeriesId(userId, seriesId)

	if _, ok := res.(entities.Series); !ok {
		return nil
	}

	dbSeasons := s.seasonRepository.FindDistinctBySeriesId(seriesId)
	var seasons []dto.SeasonDto

	for _, season := range dbSeasons {
		seasons = append(seasons, dto.SeasonDto{
			ID:       season.ID,
			ViewedAt: season.ViewedAt,
			SeriesID: season.SeriesID,
			Image:    season.Image,
			Number:   season.Number,
			Episodes: season.Episodes,
		})
	}
	return seasons
}

func (s *seasonService) GetInfosBySeasonBySeriesId(userId string, seriesId, number int) []dto.SeasonInfosDto {
	return s.seasonRepository.FindInfosBySeriesIdBySeason(userId, seriesId, number)
}

func (s *seasonService) GetDetailsSeasonsNbViewed(userId string, seriesId int) []dto.StatDto {
	return s.seasonRepository.FindDetailsSeasonsNbViewed(userId, seriesId)
}

func (s *seasonService) AddSeasonsBySeries(userId string, seriesId int, seasons dto.SeasonsCreateDto) interface{} {
	res := s.seriesRepository.FindByUserIdSeriesId(userId, seriesId)

	if _, ok := res.(entities.Series); !ok {
		return nil
	}
	for _, season := range seasons.Seasons {
		s.seasonRepository.Save(entities.Season{
			Number:   season.Number,
			Episodes: season.Episodes,
			Image:    season.Image,
			ViewedAt: seasons.ViewedAt,
			SeriesID: seriesId,
		})
	}
	return seasons
}

func (s *seasonService) GetToContinue(userId string) []dto.SeriesToContinueDto {
	apiKey := os.Getenv("API_KEY")
	series := s.seriesRepository.FindByUserIdAndWatching(userId)
	var seasons dto.SearchSeasonsDto
	var toContinue []dto.SeriesToContinueDto

	for _, userSeries := range series {

		userSeasons := s.seasonRepository.FindDistinctBySeriesId(userSeries.ID)
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

func (s *seasonService) UpdateSeason(userId string, updateDto dto.SeasonUpdateDto) interface{} {
	res := s.seasonRepository.FindById(userId, updateDto.Id)

	if season, ok := res.(entities.Season); ok {
		season.ViewedAt = updateDto.ViewedAt
		return s.seasonRepository.Save(season)
	}
	return nil
}

func (s *seasonService) DeleteSeason(userId string, seasonId int) bool {
	res := s.seasonRepository.FindById(userId, seasonId)

	if _, ok := res.(entities.Season); ok {
		return s.seasonRepository.DeleteById(seasonId)
	}
	return false
}
