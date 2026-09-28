package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	commands map[string]cliCommand
	nextURL  *string
	prevURL  *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading input:", err)
			break
		}

		userInput := scanner.Text()
		cleanedInput := cleanInput(userInput)
		firstWord := cleanedInput[0]

		value, ok := cfg.commands[firstWord]

		if ok {
			value.callback(cfg)
		} else {
			fmt.Printf("%s is not a command, type help", userInput)
			fmt.Println()
		}
	}
}

func cleanInput(text string) []string {

	lowerCaseWords := strings.ToLower(text)
	cleanWords := strings.Fields(lowerCaseWords)

	return cleanWords
}
