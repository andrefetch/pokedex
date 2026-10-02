package main

import "github.com/andrefetch/pokedex/internal/commands"

type tuiModel struct {
	cfg    *commands.Config
	input  string
	output []string
}
