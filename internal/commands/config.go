package commands

import "github.com/andrefetch/pokedex/internal/pokeapi"

type Config struct {
	PokeapiClient pokeapi.Client
	Commands      map[string]Command
	PokemonCaught map[string]pokeapi.Pokemon
	nextURL       *string
	prevURL       *string
	Pokeballs     map[string]int
}
