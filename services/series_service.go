package services

import (
	"seriesmanager-services/dto"
	"seriesmanager-services/models"
	"seriesmanager-services/repositories"
	"time"
)

type SeriesService interface {
	AddSeries(series dto.SeriesCreateDto) models.Series
	GetAll(userId string) []dto.SeriesPreviewDto
	GetByUserIdByName(userId, title string) []dto.SeriesPreviewDto
	IsDuplicateSeries(series dto.SeriesCreateDto) bool
	GetInfosBySeriesId(seriesId int) models.SeriesInfo
	DeleteByUserIdBySeriesId(userId string, seriesId int) bool
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
		Sid:           series.Sid,
		Title:         series.Title,
		Poster:        series.Poster,
		EpisodeLength: series.EpisodeLength,
		AddedAt:       time.Now(),
		UserID:        series.UserId,
	}
	return s.seriesRepository.Save(toCreate)
}

func (s *seriesService) GetAll(userId string) []dto.SeriesPreviewDto {
	var series []dto.SeriesPreviewDto
	dbSeries := s.seriesRepository.FindByUserId(userId)

	for _, s := range dbSeries {
		series = append(series, dto.SeriesPreviewDto{
			Id:            s.ID,
			Title:         s.Title,
			Poster:        s.Poster,
			EpisodeLength: s.EpisodeLength,
			Sid:           s.Sid,
		})
	}
	return series
}

func (s *seriesService) GetByUserIdByName(userId, title string) []dto.SeriesPreviewDto {
	var series []dto.SeriesPreviewDto
	dbSeries := s.seriesRepository.FindByUserIdAndName(userId, title)

	for _, s := range dbSeries {
		series = append(series, dto.SeriesPreviewDto{
			Id:            s.ID,
			Title:         s.Title,
			Poster:        s.Poster,
			EpisodeLength: s.EpisodeLength,
		})
	}
	return series
}

func (s *seriesService) IsDuplicateSeries(series dto.SeriesCreateDto) bool {
	res := s.seriesRepository.Exists(series.Sid, series.UserId)
	return res.Error == nil
}

func (s *seriesService) GetInfosBySeriesId(seriesId int) models.SeriesInfo {
	return s.seriesRepository.FindInfosBySeriesId(seriesId)
}

func (s *seriesService) DeleteByUserIdBySeriesId(userId string, seriesId int) bool {
	return s.seriesRepository.DeleteByUserBySeriesId(userId, seriesId)
}
