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
	SearchSeasonsBySid(seriesId string) dto.SearchSeasons
	SearchEpisodesBySidBySeason(seriesId string, seasonNumber string) dto.SearchEpisodes
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

func (s *searchService) SearchSeasonsBySid(sid string) dto.SearchSeasons {
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/seasons?id=%s&key=%s", sid, apiKey))
	var seasons dto.SearchSeasons

	if err := json.Unmarshal(body, &seasons); err != nil {
		panic(err.Error())
	}
	return seasons
}

func (s *searchService) SearchEpisodesBySidBySeason(sid, seasonNumber string) dto.SearchEpisodes {
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/episodes?id=%s&season=%s&key=%s", sid, seasonNumber, apiKey))
	var episodes dto.SearchEpisodes

	if err := json.Unmarshal(body, &episodes); err != nil {
		panic(err.Error())
	}
	return episodes
}
