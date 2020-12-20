package howsum

func recursive(targetSum int, numbers []int) []int {
	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	for _, n := range numbers {
		remainder := targetSum - n

		res := recursive(remainder, numbers)
		if res != nil {
			return append(res, n)
		}
	}

	return nil
}

type memo map[int][]int

func (m memo) recursive(targetSum int, numbers []int) []int {
	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	if res, ok := m[targetSum]; ok {
		return res
	}

	for _, n := range numbers {
		remainder := targetSum - n

		m[targetSum] = m.recursive(remainder, numbers)
		if m[targetSum] != nil {
			m[targetSum] = append(m[targetSum], n)
			return m[targetSum]
		}
	}

	m[targetSum] = nil
	return m[targetSum]
}

func tabulation(targetSum int, numbers []int) []int {
	targets := make([][]int, targetSum+1)

	targets[0] = []int{}

	for i := range targets {
		if targets[i] != nil {
			for _, n := range numbers {
				if i+n <= targetSum && targets[i+n] == nil {
					targets[i+n] = append(targets[i], n)
				}
			}
		}
	}

	return targets[targetSum]
}
