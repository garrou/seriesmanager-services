package services

import (
	"encoding/json"
	"fmt"
	"os"
	"services-series-manager/dto"
	"services-series-manager/helpers"

	"services-series-manager/models"
	"services-series-manager/repositories"
)

type SeriesService interface {
	AddSeries(series dto.SeriesCreateDto) models.Series
	GetAll(userId string) []models.PreviewSeries
	IsDuplicateSeries(series dto.SeriesCreateDto) bool
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
		Id:   series.Id,
		User: series.User,
	}
	return s.seriesRepository.Save(toCreate)
}

func (s *seriesService) GetAll(userId string) []models.PreviewSeries {
	apiKey := os.Getenv("API_KEY")
	dbSeries := s.seriesRepository.FindByUserId(userId)
	var series []models.PreviewSeries

	for _, s := range dbSeries {
		body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/display?id=%d&key=%s", s.Id, apiKey))
		preview := models.PreviewSeries{}

		if err := json.Unmarshal(body, &preview); err == nil {
			series = append(series, preview)
		}
	}
	return series
}

func (s *seriesService) IsDuplicateSeries(series dto.SeriesCreateDto) bool {
	res := s.seriesRepository.Exists(series.Id, series.User)
	return res.Error == nil
}
