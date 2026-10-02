package main

import (
	"fmt"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"

	"github.com/andrefetch/pokedex/internal/commands"
	"github.com/andrefetch/pokedex/internal/pokeapi"
)

func main() {
	cfg := &commands.Config{
		Commands:      commands.GetCommands(),
		PokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
		PokemonCaught: make(map[string]pokeapi.Pokemon),
		PokeballTypes: map[string]int{
			"pokeball":   10,
			"greatball":  0,
			"ultraball":  0,
			"masterball": 0,
		},
	}
	program := tea.NewProgram(tuiModel{cfg: cfg})

	if _, err := program.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
