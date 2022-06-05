package services

import (
	"github.com/google/uuid"
	"seriesmanager-services/dto"
	"seriesmanager-services/models"
	"seriesmanager-services/repositories"
	"time"
)

type SeriesService interface {
	AddSeries(series dto.SeriesCreateDto) models.Series
	GetAll(userId string) []dto.SeriesPreviewDto
	GetByTitle(userId, title string) []dto.SeriesPreviewDto
	IsDuplicateSeries(series dto.SeriesCreateDto) bool
	GetInfosBySid(sid string) models.SeriesInfo
	DeleteByUserBySid(userId, sid string) bool
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
		Id:            series.Id,
		Title:         series.Title,
		Poster:        series.Poster,
		EpisodeLength: series.EpisodeLength,
		AddedAt:       time.Now(),
		User:          series.User,
		Sid:           uuid.New().String(),
	}
	return s.seriesRepository.Save(toCreate)
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
			Sid:           s.Sid,
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
			Sid:           s.Sid,
		})
	}
	return series
}

func (s *seriesService) IsDuplicateSeries(series dto.SeriesCreateDto) bool {
	res := s.seriesRepository.Exists(series.Id, series.User)
	return res.Error == nil
}

func (s *seriesService) GetInfosBySid(sid string) models.SeriesInfo {
	return s.seriesRepository.FindInfosBySid(sid)
}

func (s *seriesService) DeleteByUserBySid(userId, sid string) bool {
	return s.seriesRepository.DeleteByUserBySid(userId, sid)
}
