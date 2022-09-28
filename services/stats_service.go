package services

import (
	"seriesmanager-services/dto"
	"seriesmanager-services/repositories"
)

func GetNbSeasonsByYears(userId string) []dto.StatDto {
	return repositories.FindNbSeasonsByYears(userId)
}

func GetTimeSeasonsByYears(userId string) []dto.StatDto {
	return repositories.FindTimeSeasonsByYears(userId)
}

func GetEpisodesByYears(userId string) []dto.StatDto {
	return repositories.FindEpisodesByYears(userId)
}

func GetTotalSeries(userId string) int64 {
	return repositories.FindTotalSeries(userId)
}

func GetTotalTime(userId string) dto.SeriesStatDto {
	return repositories.FindTotalTime(userId)
}

func GetTimeCurrentMonth(userId string) dto.SeriesStatDto {
	return repositories.FindTimeCurrentMonth(userId)
}

func GetAddedSeriesByYears(userId string) []dto.StatDto {
	return repositories.FindAddedSeriesByYears(userId)
}

func GetNbSeasonsByMonths(userId string) []dto.StatDto {
	return repositories.FindNbSeasonsByMonths(userId)
}
