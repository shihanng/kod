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

func canConstructFast(ransomNote string, magazine string) bool {
	magazineMap := make([]int, 26)

	for _, r := range magazine {
		magazineMap[r-'a']++
	}

	for _, r := range ransomNote {
		if n := magazineMap[r-'a']; n > 0 {
			magazineMap[r-'a']--
		} else {
			return false
		}
	}

	return true
}
