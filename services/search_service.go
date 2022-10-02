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
	SearchCharactersBySid(sid string) interface{}
	SearchActorInfoById(id string) interface{}
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
	var images models.Pictures
	var urls []string

	for _, series := range searchedSeries.Series {
		body = helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/pictures?id=%d&key=%s", series.Id, apiKey))
		_ = json.Unmarshal(body, &images)

		for _, u := range images.Pictures {
			urls = append(urls, u.Url)
		}
	}
	return urls
}

func (s *searchService) SearchCharactersBySid(sid string) interface{} {
	var characters models.Characters

	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/characters?id=%s&key=%s", sid, apiKey))

	if err := json.Unmarshal(body, &characters); err != nil {
		panic(err.Error())
	}
	return characters.Characters
}

func (s *searchService) SearchActorInfoById(id string) interface{} {
	var actor models.Actor
	apiKey := os.Getenv("API_KEY")
	body := helpers.HttpGet(fmt.Sprintf("https://api.betaseries.com/persons/person?id=%s&key=%s", id, apiKey))

	if err := json.Unmarshal(body, &actor); err != nil {
		panic(err.Error())
	}
	return actor.Person
}
