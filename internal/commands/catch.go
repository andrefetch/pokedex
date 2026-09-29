package commands

import (
	"fmt"
	"math/rand"

	"github.com/andrefetch/pokedex/internal/pokeapi"
)

func CommandCatch(cfg *Config, args ...string) error {
	if len(args) != 1 || args[0] == "" {
		return fmt.Errorf("usage: catch <pokemon-name>")
	}

	pokemon, err := cfg.PokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	// Higher base experience lowers the chance while keeping every Pokemon catchable.
	catchChance := 100.0 / (100.0 + float64(max(pokemon.BaseExperience, 0)))
	if rand.Float64() >= catchChance {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	if cfg.PokemonCaught == nil {
		cfg.PokemonCaught = make(map[string]pokeapi.Pokemon)
	}
	cfg.PokemonCaught[pokemon.Name] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)
	return nil
}
