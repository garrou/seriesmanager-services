package services

import (
	"encoding/json"
	"fmt"
	"os"
	"seriesmanager-services/dto"
	"seriesmanager-services/helpers"
	"seriesmanager-services/models"
)

type SearchService interface {
	Discover() dto.SearchedSeriesDto
	SearchSeriesByName(name string) dto.SearchedSeriesDto
	SearchSeasonsBySid(sid int) dto.SearchSeasonsDto
	SearchEpisodesBySidBySeason(sid, season int) dto.SearchEpisodesDto
	SearchImagesBySeriesName(name string) []string
}

type searchService struct {
}

func NewSearchService() SearchService {
	return &searchService{}
}

func (s *searchService) Discover() dto.SearchedSeriesDto {
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/discover?limit=%d&key=%s", 50, apiKey))
	var series dto.SearchedSeriesDto

	if err := json.Unmarshal(body, &series); err != nil {
		panic(err.Error())
	}
	return series
}

func (s *searchService) SearchSeriesByName(name string) dto.SearchedSeriesDto {
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/search?title=%s&key=%s", name, apiKey))
	var series dto.SearchedSeriesDto

	if err := json.Unmarshal(body, &series); err != nil {
		panic(err.Error())
	}
	return series
}

func (s *searchService) SearchSeasonsBySid(sid int) dto.SearchSeasonsDto {
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/seasons?id=%d&key=%s", sid, apiKey))
	var seasons dto.SearchSeasonsDto

	if err := json.Unmarshal(body, &seasons); err != nil {
		panic(err.Error())
	}
	return seasons
}

func (s *searchService) SearchEpisodesBySidBySeason(sid, seasonNumber int) dto.SearchEpisodesDto {
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/episodes?id=%d&season=%d&key=%s", sid, seasonNumber, apiKey))
	var episodes dto.SearchEpisodesDto

	if err := json.Unmarshal(body, &episodes); err != nil {
		panic(err.Error())
	}
	return episodes
}

func (s *searchService) SearchImagesBySeriesName(name string) []string {
	var searchedSeries dto.SearchedSeriesDto
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/search?title=%s&key=%s", name, apiKey))

	if err := json.Unmarshal(body, &searchedSeries); err != nil {
		panic(err.Error())
	}
	images := make([]models.Pictures, len(searchedSeries.Series))
	var urls []string

	for i, series := range searchedSeries.Series {
		body = helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/pictures?id=%d&key=%s", series.Id, apiKey))
		_ = json.Unmarshal(body, &images[i])

		for _, u := range images[i].Pictures {
			urls = append(urls, u.Url)
		}
	}
	return urls
}
