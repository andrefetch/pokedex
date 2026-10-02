package main

import (
	tea "charm.land/bubbletea/v2"
	"github.com/andrefetch/pokedex/internal/commands"
)

type tuiModel struct {
	cfg    *commands.Config
	input  string
	output []string
}

func (m tuiModel) Init() tea.Cmd {
	return nil
}

func (m tuiModel) View() tea.View {
	return tea.NewView(
		"POKEDEX\n" +
			"Type help for commands.\n\n" +
			"Pokedex > " + m.input,
	)
}

func (m tuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if msg.Keystroke() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	return m, nil
}
