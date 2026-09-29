package commands

import "github.com/andrefetch/pokedex/internal/pokeapi"

type Config struct {
	PokeapiClient pokeapi.Client
	Commands      map[string]Command
	nextURL       *string
	prevURL       *string
}
