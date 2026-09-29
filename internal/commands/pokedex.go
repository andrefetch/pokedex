package commands

import (
	"fmt"
)

func CommandPokedex(cfg *Config, args ...string) error {

	fmt.Println("Your pokedex:")

	for _, pokemon := range cfg.PokemonCaught {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil

}
