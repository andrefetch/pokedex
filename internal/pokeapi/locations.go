package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
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

func (c *Client) ListLocations(pageURL *string) (LocationsResponse, error) {
	fullURL := baseURL + "/location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}

	var locations LocationsResponse
	if body, ok := c.cache.Get(fullURL); ok {
		err := json.Unmarshal(body, &locations)
		return locations, err
	}

	resp, err := c.httpClient.Get(fullURL)
	if err != nil {
		return locations, fmt.Errorf("retrieve locations: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return locations, fmt.Errorf("retrieve locations: unexpected status %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return locations, fmt.Errorf("read locations: %w", err)
	}
	if err := json.Unmarshal(body, &locations); err != nil {
		return locations, fmt.Errorf("decode locations: %w", err)
	}

	c.cache.Add(fullURL, body)
	return locations, nil
}
