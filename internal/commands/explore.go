package commands

import "fmt"

func CommandExplore(cfg *Config, args ...string) error {
	if len(args) != 1 || args[0] == "" {
		return fmt.Errorf("usage: explore <area-name>")
	}

	area, err := cfg.PokeapiClient.GetLocationArea(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", area.Name)
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
