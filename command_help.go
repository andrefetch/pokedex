package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n ")

	for _, value := range registry {
		fmt.Println(value.name + ": " + value.description)
	}

	return nil
}
