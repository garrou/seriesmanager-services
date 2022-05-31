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
}

type searchService struct {
}

func NewSearchService() SearchService {
	return &searchService{}
}

func (s *searchService) Discover() models.SearchSeries {
	apiKey := os.Getenv("API_KEY")
	series := models.SearchSeries{}
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/discover?limit=%d&key=%s", 30, apiKey))

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
