package services

import (
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
	"seriesmanager-services/repositories"
	"time"
)

func AddSeries(series dto.SeriesCreateDto) dto.SeriesDto {
	toCreate := entities.Series{
		Sid:           series.Sid,
		Title:         series.Title,
		Poster:        series.Poster,
		EpisodeLength: series.EpisodeLength,
		AddedAt:       time.Now(),
		UserID:        series.UserId,
	}
	created := repositories.SaveSeries(toCreate)

	return dto.SeriesDto{
		ID:            created.ID,
		Title:         created.Title,
		Poster:        created.Poster,
		EpisodeLength: created.EpisodeLength,
		Sid:           created.Sid,
		AddedAt:       created.AddedAt,
	}
}

func GetAllSeries(userId string) []dto.SeriesPreviewDto {
	var series []dto.SeriesPreviewDto
	dbSeries := repositories.FindByUserId(userId)

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

func GetByUserIdByName(userId, title string) []dto.SeriesPreviewDto {
	var series []dto.SeriesPreviewDto
	dbSeries := repositories.FindByUserIdAndName(userId, title)

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

func IsDuplicateSeries(series dto.SeriesCreateDto) bool {
	res := repositories.SeriesExists(series.Sid, series.UserId)
	return res.Error == nil
}

func GetInfosBySeriesId(userId string, seriesId int) dto.SeriesInfoDto {
	return repositories.FindInfosBySeriesId(userId, seriesId)
}

func DeleteByUserIdBySeriesId(userId string, seriesId int) bool {
	return repositories.DeleteByUserBySeriesId(userId, seriesId)
}

func UpdateWatching(userId string, seriesId int) interface{} {
	res := repositories.FindByUserIdSeriesId(userId, seriesId)

	if series, ok := res.(entities.Series); ok {
		series.IsWatching = !series.IsWatching
		return repositories.SaveSeries(series)
	}
	return nil
}
