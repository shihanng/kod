package note

func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := make(map[rune]int, len(magazine))

	for _, r := range magazine {
		magazineMap[r]++
	}

	for _, r := range ransomNote {
		if n, ok := magazineMap[r]; ok && n > 0 {
			magazineMap[r]--
		} else {
			return false
		}
	}

	return true
}
