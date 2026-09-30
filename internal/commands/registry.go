package commands

type Command struct {
	Name        string
	Description string
	Callback    func(*Config, ...string) error
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "Displays this message",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"map": {
			Name:        "map",
			Description: "Displays a list of map locations you can explore.",
			Callback:    CommandMapf,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays a list of map locations 20 results backwards.",
			Callback:    CommandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "explore <area-name>: Lists Pokemon found in a map location.",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "catch <pokemon-name>: attempts to catch a pokemon in that desired area.",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "inspect <pokemon-name>: Displays details of a caught Pokemon.",
			Callback:    CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "displays all pokemon caught, in your pokedex.",
			Callback:    CommandPokedex,
		},
	}
}
