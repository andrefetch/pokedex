package commands

import "fmt"

func CommandMapf(cfg *Config) error {
	return displayLocations(cfg, cfg.nextURL)
}

func CommandMapb(cfg *Config) error {
	if cfg.prevURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	return displayLocations(cfg, cfg.prevURL)
}

func displayLocations(cfg *Config, pageURL *string) error {
	locations, err := cfg.PokeapiClient.ListLocations(pageURL)
	if err != nil {
		return err
	}

	cfg.nextURL = locations.Next
	cfg.prevURL = locations.Previous
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
