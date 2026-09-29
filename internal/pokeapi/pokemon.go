package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	fullURL := baseURL + "/pokemon/" + url.PathEscape(name)
	var pokemon Pokemon
	if body, ok := c.cache.Get(fullURL); ok {
		err := json.Unmarshal(body, &pokemon)
		return pokemon, err
	}

	resp, err := c.httpClient.Get(fullURL)
	if err != nil {
		return pokemon, fmt.Errorf("retrieve pokemon: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return pokemon, fmt.Errorf("retrieve pokemon: unexpected status %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokemon, fmt.Errorf("read pokemon: %w", err)
	}
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return pokemon, fmt.Errorf("decode pokemon: %w", err)
	}
	c.cache.Add(fullURL, body)
	return pokemon, nil
}
