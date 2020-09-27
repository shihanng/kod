// Package scrabble provides helper function to compute scrabble score of a word.
package scrabble

import (
	"strings"
)

var scores = map[rune]int{
	'A': 1,
	'E': 1,
	'I': 1,
	'O': 1,
	'U': 1,
	'L': 1,
	'N': 1,
	'R': 1,
	'S': 1,
	'T': 1,
	'D': 2,
	'G': 2,
	'B': 3,
	'C': 3,
	'M': 3,
	'P': 3,
	'F': 4,
	'H': 4,
	'V': 4,
	'W': 4,
	'Y': 4,
	'K': 5,
	'J': 8,
	'X': 8,
	'Q': 10,
	'Z': 10,
}

// Score computes scrabble score of a given word.
func Score(word string) int {
	normalizedWord := strings.ToUpper(word)

	var s int

	for _, r := range normalizedWord {
		s += scores[r]
	}

	return s
}
