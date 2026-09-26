package main

import (
	"strings"
)

func cleanInput(text string) []string {

	lowerCaseWords := strings.ToLower(text)
	cleanWords := strings.Fields(lowerCaseWords)

	return cleanWords
}
