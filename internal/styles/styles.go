package styles

import "charm.land/lipgloss/v2"

var (
	PokedexPromptStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#E63946")).Bold(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF5555")).Bold(true)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#50FA7B")).Bold(true)

	GreyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#464646")).Bold(true)

	InfoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#9edbe0")).Bold(true)
)

var Card = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#E63946")).
	Padding(1, 2)
