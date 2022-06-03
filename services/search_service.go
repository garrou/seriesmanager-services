package services

import (
	"encoding/json"
	"fmt"
	"os"
	"seriesmanager-services/helpers"
	"seriesmanager-services/models"
)

type SearchService interface {
	Discover() models.SearchSeries
	SearchSeriesByName(name string) models.SearchSeries
	SearchSeriesById(seriesId string) models.DetailsSeries
	SearchSeasonsBySeriesId(seriesId string) models.SearchSeasons
}

type searchService struct {
}

func NewSearchService() SearchService {
	return &searchService{}
}

func (s *searchService) Discover() models.SearchSeries {
	apiKey := os.Getenv("API_KEY")
	series := models.SearchSeries{}
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/discover?limit=%d&key=%s", 20, apiKey))

	if err := json.Unmarshal(body, &series); err != nil {
		panic(err.Error())
	}
	return series
}

func (s *searchService) SearchSeriesByName(name string) models.SearchSeries {
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/search?title=%s&key=%s", name, apiKey))
	series := models.SearchSeries{}

	if err := json.Unmarshal(body, &series); err != nil {
		panic(err.Error())
	}
	return series
}

func (s *searchService) SearchSeriesById(seriesId string) models.DetailsSeries {
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/display?id=%s&key=%s", seriesId, apiKey))
	series := models.DetailsSeries{}

	if err := json.Unmarshal(body, &series); err != nil {
		panic(err.Error())
	}
	return series
}

func (s *searchService) SearchSeasonsBySeriesId(seriesId string) models.SearchSeasons {
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/seasons?id=%s&key=%s", seriesId, apiKey))
	var seasons models.SearchSeasons

	if err := json.Unmarshal(body, &seasons); err != nil {
		panic(err.Error())
	}
	return seasons
}
