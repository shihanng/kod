package canconstruct

import "strings"

type memo map[string]bool

func (m memo) recursive(target string, elements []string) bool {
	if res, ok := m[target]; ok {
		return res
	}

	if target == "" {
		return true
	}

	for _, elem := range elements {
		if !strings.HasPrefix(target, elem) {
			continue
		}

		if m.recursive(strings.TrimPrefix(target, elem), elements) {
			m[target] = true
			return true
		}
	}

	m[target] = false
	return false
}

func tabulation(target string, elements []string) bool {
	results := make([]bool, len(target)+1)
	results[0] = true

	for i := range results {
		if !results[i] {
			continue
		}
		for _, elem := range elements {
			if strings.HasPrefix(target[i:], elem) {
				results[len(elem)+i] = true
			}
		}
	}

	return results[len(target)]
}
