package valid_parent

func minAddToMakeValid(S string) int {
	if S == "" {
		return 0
	}

	leftCount := 0
	rightCount := 0

	for _, r := range S {
		if r == '(' {
			leftCount++
		} else {
			if leftCount > 0 {
				leftCount--
			} else {
				rightCount++
			}
		}
	}

	return leftCount + rightCount
}
