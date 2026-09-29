package main

import (
	"time"

	"github.com/andrefetch/pokedex/internal/commands"
	"github.com/andrefetch/pokedex/internal/pokeapi"
)

func main() {
	cfg := &commands.Config{
		Commands:      commands.GetCommands(),
		PokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
	}
	startRepl(cfg)
}
