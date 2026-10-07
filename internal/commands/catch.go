package commands

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/andrefetch/pokedex/internal/pokeapi"
	"github.com/andrefetch/pokedex/internal/styles"
)

func CommandCatch(cfg *Config, args ...string) error {
	if len(args) != 2 || args[0] == "" || args[1] == "" {
		return fmt.Errorf("usage: catch <pokemon-name> <ball-type>")
	}

	pokemonName := args[0]
	pokeballType := args[1]

	ballTypes := map[string]Pokeball{
		"pokeball":   StandardPokeball,
		"greatball":  GreatBall,
		"ultraball":  UltraBall,
		"masterball": MasterBall,
	}

	ball, ok := ballTypes[pokeballType]
	if !ok {
		return fmt.Errorf("%s is an unsupported ball type", pokeballType)
	}

	if cfg.PokeballTypes[pokeballType] <= 0 {
		return fmt.Errorf("you need at least 1 %s to attempt a catch", pokeballType)
	}

	pokemon, err := cfg.PokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	randShake := rand.Intn(4) + 1

	cfg.PokeballTypes[pokeballType]--
	message := fmt.Sprintf("Throwing a %s at %s...\n", ball.Name, pokemon.Name)
	fmt.Println(styles.InfoStyle.Render(message))

	for i := 0; i < randShake; i++ {
		time.Sleep(1 * time.Second)
		message := fmt.Sprintf("%s shook...", ball.Name)
		fmt.Println(styles.GreyStyle.Render(message))
	}

	time.Sleep(2 * time.Second)

	// Higher base experience lowers the catch chance.
	baseChance := 100.0 / (100.0 + float64(max(pokemon.BaseExperience, 0)))
	catchChance := min(1.0, baseChance*ball.CatchMultiplier)

	if !ball.GaurenteedCatch && rand.Float64() >= catchChance {
		message := fmt.Sprintf("%s escaped!\n", pokemon.Name)
		fmt.Println(styles.ErrorStyle.Render(message))
		return nil
	}

	if cfg.PokemonCaught == nil {
		cfg.PokemonCaught = make(map[string]pokeapi.Pokemon)
	}

	cfg.PokemonCaught[pokemon.Name] = pokemon
	caughtMessage := fmt.Sprintf("%s was caught!\n", pokemon.Name)
	fmt.Println(styles.SuccessStyle.Render(caughtMessage))
	return nil
}
