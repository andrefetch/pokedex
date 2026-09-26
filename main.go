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
		firstWordofCmd := cleanedInput[0]
		fmt.Println("Your command was:", firstWordofCmd)
	}

}
