package countconstruct

import "strings"

type memo map[string]int

func (m memo) recursive(target string, elements []string) int {
	if res, ok := m[target]; ok {
		return res
	}

	if target == "" {
		return 1
	}

	var count int
	for _, elem := range elements {
		if !strings.HasPrefix(target, elem) {
			continue
		}

		if n := m.recursive(strings.TrimPrefix(target, elem), elements); n > 0 {
			count += n
		}
	}

	m[target] = count
	return count
}

func tabulation(target string, elements []string) int {
	results := make([]int, len(target)+1)
	results[0] = 1

	for i := range results {
		if results[i] == 0 {
			continue
		}
		for _, elem := range elements {
			if strings.HasPrefix(target[i:], elem) {
				results[len(elem)+i] += results[i]
			}
		}
	}

	return results[len(target)]
}
