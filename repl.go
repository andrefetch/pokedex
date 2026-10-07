package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/andrefetch/pokedex/internal/commands"
	"github.com/andrefetch/pokedex/internal/styles"
)

func startRepl(cfg *commands.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(styles.PokedexPromptStyle.Render("Pokedex > "))

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				message := fmt.Sprintf("error reading input: %v", err)
				fmt.Fprintln(os.Stderr, styles.ErrorStyle.Render(message))
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
			message := fmt.Sprintf("%s is not a command, type help", userInput)
			fmt.Fprintln(os.Stderr, styles.ErrorStyle.Render(message))
			fmt.Println()
		}
	}
}

func cleanInput(text string) []string {

	lowerCaseWords := strings.ToLower(text)
	cleanWords := strings.Fields(lowerCaseWords)

	return cleanWords
}
