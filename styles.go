package main

import "charm.land/lipgloss/v2"

var (
	pokedexPromptStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#E63946")).Bold(true)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF5555")).Bold(true)

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#50FA7B"))
)
