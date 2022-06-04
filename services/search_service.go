package services

import (
	"encoding/json"
	"fmt"
	"os"
	"seriesmanager-services/dto"
	"seriesmanager-services/helpers"
)

type SearchService interface {
	Discover() dto.SearchSeries
	SearchSeriesByName(name string) dto.SearchSeries
	SearchSeriesById(seriesId string) dto.DetailsSeries
	SearchSeasonsBySeriesId(seriesId string) dto.SearchSeasons
	SearchEpisodesBySeriesIdBySeason(seriesId string, seasonNumber string) dto.SearchEpisodes
}

type searchService struct {
}

func NewSearchService() SearchService {
	return &searchService{}
}

func (s *searchService) Discover() dto.SearchSeries {
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/discover?limit=%d&key=%s", 20, apiKey))
	var series dto.SearchSeries

	if err := json.Unmarshal(body, &series); err != nil {
		panic(err.Error())
	}
	return series
}

func (s *searchService) SearchSeriesByName(name string) dto.SearchSeries {
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/search?title=%s&key=%s", name, apiKey))
	var series dto.SearchSeries

	if err := json.Unmarshal(body, &series); err != nil {
		panic(err.Error())
	}
	return series
}

func (s *searchService) SearchSeriesById(seriesId string) dto.DetailsSeries {
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/display?id=%s&key=%s", seriesId, apiKey))
	var series dto.DetailsSeries

	if err := json.Unmarshal(body, &series); err != nil {
		panic(err.Error())
	}
	return series
}

func (s *searchService) SearchSeasonsBySeriesId(seriesId string) dto.SearchSeasons {
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/seasons?id=%s&key=%s", seriesId, apiKey))
	var seasons dto.SearchSeasons

	if err := json.Unmarshal(body, &seasons); err != nil {
		panic(err.Error())
	}
	return seasons
}

func (s *searchService) SearchEpisodesBySeriesIdBySeason(seriesId string, seasonNumber string) dto.SearchEpisodes {
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/episodes?id=%s&season=%s&key=%s", seriesId, seasonNumber, apiKey))
	var episodes dto.SearchEpisodes

	if err := json.Unmarshal(body, &episodes); err != nil {
		panic(err.Error())
	}
	return episodes
}
