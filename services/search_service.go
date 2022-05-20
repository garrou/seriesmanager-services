package services

import (
	"encoding/json"
	"fmt"
	"os"
	"services-series-manager/helpers"
	"services-series-manager/models"
)

type SearchService interface {
	SearchSeriesByName(name string) models.SearchSeries
}

type searchService struct {
}

func NewSearchService() SearchService {
	return &searchService{}
}

func (s *searchService) SearchSeriesByName(name string) models.SearchSeries {
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/search?title=%s&key=%s", name, apiKey))
	previewSeries := models.SearchSeries{}

	if err := json.Unmarshal(body, &previewSeries); err != nil {
		panic(err.Error())
	}
	return previewSeries
}
