package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocationsResponse struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func commandMapf(cfg *config) error {

	const baseURL = "https://pokeapi.co/api/v2/location-area"

	fullURL := baseURL

	if cfg.nextURL != nil {
		fullURL = *cfg.nextURL
	}

	req, err := http.Get(fullURL)
	if err != nil {
		fmt.Errorf("Error occured while retrieving data", err)
		return err
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Errorf("Error occured while retrieving data", err)
		return err
	}

	defer req.Body.Close()

	var locationsResponse LocationsResponse
	if err := json.Unmarshal(body, &locationsResponse); err != nil {
		log.Fatalf("Unmarshal error: %v", err)
	}

	cfg.nextURL = locationsResponse.Next
	cfg.prevURL = locationsResponse.Previous

	for _, loc := range locationsResponse.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
