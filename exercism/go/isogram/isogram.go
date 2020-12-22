package isogram

import "strings"

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
