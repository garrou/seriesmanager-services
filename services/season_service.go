package services

import (
	"encoding/json"
	"fmt"
	"os"
	"seriesmanager-services/dto"
	"seriesmanager-services/entities"
	"seriesmanager-services/helpers"
	"seriesmanager-services/repositories"
	"sync"
)

func GetDistinctBySeriesId(userId string, seriesId int) []dto.SeasonDto {
	res := repositories.FindByUserIdSeriesId(userId, seriesId)

	if _, ok := res.(entities.Series); !ok {
		return nil
	}

	dbSeasons := repositories.FindDistinctBySeriesId(seriesId)
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

func GetInfosBySeasonBySeriesId(userId string, seriesId, number int) []dto.SeasonInfosDto {
	return repositories.FindInfosBySeriesIdBySeason(userId, seriesId, number)
}

func GetDetailsSeasonsNbViewed(userId string, seriesId int) []dto.StatDto {
	return repositories.FindDetailsSeasonsNbViewed(userId, seriesId)
}

func AddSeasonsBySeries(userId string, seriesId int, seasons dto.SeasonsCreateDto) interface{} {
	res := repositories.FindByUserIdSeriesId(userId, seriesId)

	if _, ok := res.(entities.Series); !ok {
		return nil
	}
	for _, season := range seasons.Seasons {
		repositories.SaveSeason(entities.Season{
			Number:   season.Number,
			Episodes: season.Episodes,
			Image:    season.Image,
			ViewedAt: seasons.ViewedAt,
			SeriesID: seriesId,
		})
	}
	return seasons
}

func GetToContinue(userId string) []dto.SeriesToContinueDto {
	var seasons dto.SearchSeasonsDto
	var toContinue []dto.SeriesToContinueDto
	var wg sync.WaitGroup
	series := repositories.FindByUserIdAndWatching(userId)
	apiKey := os.Getenv("API_KEY")

	for _, userSeries := range series {
		wg.Add(1)
		go func(series entities.Series) {
			defer wg.Done()
			userSeasons := repositories.FindDistinctBySeriesId(series.ID)
			body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/seasons?id=%d&key=%s", series.Sid, apiKey))

			if err := json.Unmarshal(body, &seasons); err != nil {
				panic(err.Error())
			}
			diff := len(seasons.Seasons) - len(userSeasons)

			if diff > 0 {
				toContinue = append(toContinue, dto.SeriesToContinueDto{
					Id:            series.ID,
					Title:         series.Title,
					Poster:        series.Poster,
					EpisodeLength: series.EpisodeLength,
					Sid:           series.Sid,
					NbMissing:     diff,
				})
			}
		}(userSeries)
	}
	wg.Wait()
	return toContinue
}

func UpdateSeason(userId string, updateDto dto.SeasonUpdateDto) interface{} {
	res := repositories.FindSeasonById(userId, updateDto.Id)

	if season, ok := res.(entities.Season); ok {
		season.ViewedAt = updateDto.ViewedAt
		return repositories.SaveSeason(season)
	}
	return nil
}

func DeleteSeason(userId string, seasonId int) bool {
	res := repositories.FindSeasonById(userId, seasonId)

	if _, ok := res.(entities.Season); ok {
		return repositories.DeleteSeasonById(seasonId)
	}
	return false
}
