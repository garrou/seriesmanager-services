package services

import (
	"encoding/json"
	"fmt"
	"os"
	"seriesmanager-services/dto"
	"seriesmanager-services/helpers"
)

type SearchService interface {
	Discover() dto.SearchedSeries
	SearchSeriesByName(name string) dto.SearchedSeries
	SearchSeasonsBySid(seriesId string) dto.SearchSeasons
	SearchEpisodesBySidBySeason(seriesId string, seasonNumber string) dto.SearchEpisodes
	SearchImagesBySeriesName(name string) []string
}

type searchService struct {
}

func NewSearchService() SearchService {
	return &searchService{}
}

func (s *searchService) Discover() dto.SearchedSeries {
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/discover?limit=%d&key=%s", 20, apiKey))
	var series dto.SearchedSeries

	if err := json.Unmarshal(body, &series); err != nil {
		panic(err.Error())
	}
	return series
}

func (s *searchService) SearchSeriesByName(name string) dto.SearchedSeries {
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/search?title=%s&key=%s", name, apiKey))
	var series dto.SearchedSeries

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

func (s *searchService) SearchImagesBySeriesName(name string) []string {
	var searchedSeries dto.SearchedSeries
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/search?title=%s&key=%s", name, apiKey))

	if err := json.Unmarshal(body, &searchedSeries); err != nil {
		panic(err.Error())
	}
	images := make([]dto.Pictures, len(searchedSeries.Series))
	var urls []string

	for i, series := range searchedSeries.Series {
		body = helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/pictures?id=%d&key=%s", series.Id, apiKey))
		json.Unmarshal(body, &images[i])

		for _, u := range images[i].Pictures {
			urls = append(urls, u.Url)
		}
	}
	return urls
}
