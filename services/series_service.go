package services

import (
	"services-series-manager/dto"

	"services-series-manager/models"
	"services-series-manager/repositories"
)

type SeriesService interface {
	AddSeries(user dto.SeriesCreateDto) models.Series
	GetAll(userId string) []models.Series
}

type seriesService struct {
	seriesRepository repositories.SeriesRepository
}

func NewSeriesService(seriesRepository repositories.SeriesRepository) SeriesService {
	return &seriesService{
		seriesRepository: seriesRepository,
	}
}

func (s *seriesService) AddSeries(series dto.SeriesCreateDto) models.Series {
	toCreate := models.Series{
		Id:     series.Id,
		FkUser: series.FkUser,
	}
	return s.seriesRepository.Save(toCreate)
}

func (s *seriesService) GetAll(userId string) []models.Series {
	return s.seriesRepository.FindByUserId(userId)
}
