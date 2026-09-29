package commands

import (
	"errors"
	"fmt"
)

func CommandInspect(cfg *Config, args ...string) error {
	if len(args) != 1 || args[0] == "" {
		return errors.New("you must provide a pokemon name")
	}

	pokemonData, exists := cfg.PokemonCaught[args[0]]
	if !exists {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemonData.Name)
	fmt.Printf("Height: %d\n", pokemonData.Height)
	fmt.Printf("Weight: %d\n", pokemonData.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonData.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemonData.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}
