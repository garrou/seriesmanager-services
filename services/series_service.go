package services

import (
	"seriesmanager-services/dto"
	"seriesmanager-services/models"
	"seriesmanager-services/repositories"
)

type SeriesService interface {
	AddSeries(series dto.SeriesCreateDto) dto.SeriesPreviewDto
	GetAll(userId string) []dto.SeriesPreviewDto
	GetByTitle(userId, title string) []dto.SeriesPreviewDto
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

func (s *seriesService) AddSeries(series dto.SeriesCreateDto) dto.SeriesPreviewDto {
	toCreate := models.Series{
		Id:            series.Id,
		Title:         series.Title,
		Poster:        series.Poster,
		EpisodeLength: series.EpisodeLength,
		User:          series.User,
	}
	s.seriesRepository.Save(toCreate)
	return dto.SeriesPreviewDto{Id: series.Id, Title: series.Title, Poster: series.Poster}
}

func (s *seriesService) GetAll(userId string) []dto.SeriesPreviewDto {
	var series []dto.SeriesPreviewDto
	dbSeries := s.seriesRepository.FindByUserId(userId)

	for _, s := range dbSeries {
		series = append(series, dto.SeriesPreviewDto{
			Id:            s.Id,
			Title:         s.Title,
			Poster:        s.Poster,
			EpisodeLength: s.EpisodeLength,
		})
	}
	return series
}

func (s *seriesService) GetByTitle(userId, title string) []dto.SeriesPreviewDto {
	var series []dto.SeriesPreviewDto
	dbSeries := s.seriesRepository.FindByUserIdAndTitle(userId, title)

	for _, s := range dbSeries {
		series = append(series, dto.SeriesPreviewDto{
			Id:            s.Id,
			Title:         s.Title,
			Poster:        s.Poster,
			EpisodeLength: s.EpisodeLength,
		})
	}
	return series
}

func (s *seriesService) IsDuplicateSeries(series dto.SeriesCreateDto) bool {
	res := s.seriesRepository.Exists(series.Id, series.User)
	return res.Error == nil
}
