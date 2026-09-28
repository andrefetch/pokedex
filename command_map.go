package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
		return fmt.Errorf("Error occured while retrieving data: %v", err)
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return fmt.Errorf("Error occured while retrieving data: %v", err)
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

func commandMapb(cfg *config) error {

	const baseURL = "https://pokeapi.co/api/v2/location-area"

	fullURL := baseURL

	if cfg.prevURL != nil {
		fullURL = *cfg.prevURL
	} else {
		fmt.Println("you're on the first page")
		os.Exit(1)
	}

	req, err := http.Get(fullURL)
	if err != nil {
		return fmt.Errorf("Error occured while retrieving data: %v", err)
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return fmt.Errorf("Error occured while retrieving data: %v", err)
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
