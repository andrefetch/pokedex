package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/andrefetch/pokedex/internal/commands"
)

func startRepl(cfg *commands.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(pokedexPromptStyle.Render("Pokedex > "))

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "error reading input:", err)
			}
			break
		}

		userInput := scanner.Text()
		cleanedInput := cleanInput(userInput)
		if len(cleanedInput) == 0 {
			continue
		}
		firstWord := cleanedInput[0]

		value, ok := cfg.Commands[firstWord]

		if ok {
			if err := value.Callback(cfg, cleanedInput[1:]...); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
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
