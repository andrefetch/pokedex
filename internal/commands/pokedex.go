package commands

import (
	"fmt"

	"github.com/andrefetch/pokedex/internal/styles"
)

func CommandPokedex(cfg *Config, args ...string) error {

	fmt.Println("Your pokedex:")

	for _, pokemon := range cfg.PokemonCaught {
		content := fmt.Sprintf(
			"Name: %s\nHeight: %d\nWeight: %d\n",
			pokemon.Name,
			pokemon.Height,
			pokemon.Weight,
		)

		fmt.Println(styles.Card.Render(content))
	}
	return nil

}
