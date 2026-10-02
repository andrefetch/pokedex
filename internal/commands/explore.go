package commands

import (
	"fmt"
	"math/rand"
)

func CommandExplore(cfg *Config, args ...string) error {
	if len(args) != 1 || args[0] == "" {
		return fmt.Errorf("usage: explore <area-name>")
	}

	area, err := cfg.PokeapiClient.GetLocationArea(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", area.Name)

	// Pokeball rolling chance (chances are in pokeball.go)
	ballOptions := []struct {
		key  string
		ball Pokeball
	}{
		{"pokeball", StandardPokeball},
		{"greatball", GreatBall},
		{"ultraball", UltraBall},
		{"masterball", MasterBall},
	}

	roll := rand.Float64()
	threshold := 0.0

	for _, option := range ballOptions {
		threshold += option.ball.SpawnChance

		if roll < float64(threshold) {
			cfg.PokeballTypes[option.key]++
			fmt.Printf("You have found a %s!\n", option.ball.Name)
			break
		}
	}

	if len(area.PokemonEncounters) == 0 {
		fmt.Println("No Pokemon found.")
		return nil
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range area.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
