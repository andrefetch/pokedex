package commands

import (
	"fmt"
)

func CommandHelp(cfg *Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n ")

	for _, value := range cfg.Commands {
		fmt.Println(value.Name + ": " + value.Description)
	}

	return nil
}
