package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type LocationAreaResponse struct {
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
	} `json:"pokemon"`
}

func (c *Client) GetLocationArea(name string) (LocationAreaResponse, error) {
	fullURL := baseURL + "/location-area/" + url.PathEscape(name)
	var area LocationAreaResponse
	if body, ok := c.cache.Get(fullURL); ok {
		err := json.Unmarshal(body, &area)
		return area, err
	}

	resp, err := c.httpClient.Get(fullURL)
	if err != nil {
		return area, fmt.Errorf("retrieve location area: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return area, fmt.Errorf("retrieve location area: unexpected status %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return area, fmt.Errorf("read location area: %w", err)
	}
	if err := json.Unmarshal(body, &area); err != nil {
		return area, fmt.Errorf("decode location area: %w", err)
	}
	c.cache.Add(fullURL, body)
	return area, nil
}
