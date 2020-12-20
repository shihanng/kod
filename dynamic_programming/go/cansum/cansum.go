package cansum

func recursive(targetSum int, numbers []int) bool {
	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	for _, n := range numbers {
		remainder := targetSum - n
		if recursive(remainder, numbers) {
			return true
		}
	}

	return false
}

type memo map[int]bool

func (m memo) recursive(targetSum int, numbers []int) bool {
	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	if res, ok := m[targetSum]; ok {
		return res
	}

	for _, n := range numbers {
		remainder := targetSum - n
		m[targetSum] = m.recursive(remainder, numbers)
		if m[targetSum] {
			return true
		}
	}

	m[targetSum] = false

	return false
}

func tabulation(targetSum int, numbers []int) bool {
	targets := make([]bool, targetSum+1)
	targets[0] = true

	for i := range targets {
		if targets[i] {
			for _, n := range numbers {
				if i+n <= targetSum {
					targets[i+n] = true
				}
			}
		}
	}

	return targets[targetSum]
}
