package isogram

import "strings"

// IsIsogram checks if the given input does not contain repeating letter.
// Spaces and hyphens are allowed to appear multiple times.
func IsIsogram(input string) bool {
	lowerInput := strings.ToLower(input)

	founds := make(map[rune]bool, len(lowerInput))

	for _, c := range lowerInput {
		if c == '-' || c == ' ' {
			continue
		}

		if founds[c] {
			return false
		}

		founds[c] = true
	}

	return true
}
