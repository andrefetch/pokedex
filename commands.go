package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays a list of map locations you can explore.",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays a list of map locations 20 results backwards.",
			callback:    commandMapb,
		},
	}
}
