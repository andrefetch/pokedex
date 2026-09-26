package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "     hEy   newWord     cRaZyUpPeRcAsInG",
			expected: []string{"hey", "newword", "crazyuppercasing"},
		},
		{
			input:    "pokedex    is super cool   man!",
			expected: []string{"pokedex", "is", "super", "cool", "man!"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			t.Errorf("Length of actual: %v, does not match expected length of: %v", len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Actual word: %v, does not match expected word: %v", word, expectedWord)
			}
		}
	}

}
