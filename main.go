package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

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

		value, ok := registry[firstWord]

		if ok {
			value.callback()
		} else {
			fmt.Println("Command not found.")
		}
	}

}
