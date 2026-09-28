package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n ")

	for _, value := range cfg.commands {
		fmt.Println(value.name + ": " + value.description)
	}

	return nil
}
