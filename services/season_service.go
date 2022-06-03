package services

import (
	"seriesmanager-services/dto"
	"seriesmanager-services/models"
	"seriesmanager-services/repositories"
	"strconv"
	"time"
)

const LAYOUT = "2000-03-14"

type SeasonService interface {
	AddSeason(season dto.SeasonCreateDto) dto.SeasonCreateDto
	GetBySeriesId(id string) []dto.SeasonDto
}

type seasonService struct {
	seasonRepository repositories.SeasonRepository
}

func NewSeasonService(seasonRepository repositories.SeasonRepository) SeasonService {
	return &seasonService{
		seasonRepository: seasonRepository,
	}
}

func (s *seasonService) AddSeason(season dto.SeasonCreateDto) dto.SeasonCreateDto {
	start, _ := time.Parse(LAYOUT, season.StartedAt)
	finish, _ := time.Parse(LAYOUT, season.FinishedAt)
	toCreate := models.Season{
		Number:     season.Number,
		Episodes:   season.Episodes,
		Image:      season.Image,
		StartedAt:  start,
		FinishedAt: finish,
		Series:     season.Series,
	}
	s.seasonRepository.Save(toCreate)
	return season
}

func (s *seasonService) GetBySeriesId(id string) []dto.SeasonDto {
	var seasons []dto.SeasonDto
	seriesId, _ := strconv.Atoi(id)
	res := s.seasonRepository.FindBySeriesId(seriesId)

	for _, s := range res {
		seasons = append(seasons, dto.SeasonDto{
			Id:         s.Id,
			Number:     s.Number,
			Episodes:   s.Episodes,
			Image:      s.Image,
			StartedAt:  s.StartedAt.Format(LAYOUT),
			FinishedAt: s.FinishedAt.Format(LAYOUT),
		})
	}
	return seasons
}
