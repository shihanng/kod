package bestsum

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

	var shortest []int

	for _, n := range numbers {
		remainder := targetSum - n

		res := m.recursive(remainder, numbers)
		if res == nil {
			continue
		}

		res = append(res, n)

		if len(shortest) == 0 || len(res) < len(shortest) {
			shortest = res
		}
	}

	m[targetSum] = shortest
	return shortest
}

func tabulation(targetSum int, numbers []int) []int {
	targets := make([][]int, targetSum+1)

	targets[0] = []int{}

	for i := range targets {
		if targets[i] != nil {
			for _, n := range numbers {
				newSum := append(targets[i], n)
				if i+n <= targetSum && (targets[i+n] == nil || len(newSum) <= len(targets[i+n])) {
					targets[i+n] = newSum
				}
			}
		}
	}

	return targets[targetSum]
}
