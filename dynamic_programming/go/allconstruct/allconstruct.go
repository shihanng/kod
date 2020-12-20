package allconstruct

import (
	"strings"
)

type memo map[string][][]string

func (m memo) recursive(target string, elements []string) [][]string {
	if res, ok := m[target]; ok {
		return res
	}

	if target == "" {
		return [][]string{}
	}

	var all [][]string

	for _, elem := range elements {
		if !strings.HasPrefix(target, elem) {
			continue
		}

		res := m.recursive(strings.TrimPrefix(target, elem), elements)
		if res == nil {
			continue
		}

		tmp := []string{elem}
		if len(res) == 0 {
			all = append(all, tmp)
		} else {
			for _, r := range res {
				withElem := append(tmp, r...)
				all = append(all, withElem)
			}
		}

	}

	m[target] = all

	return all
}

func tabulation(target string, elements []string) [][]string {
	results := make([][][]string, len(target)+1)

	results[0] = [][]string{}

	for i := range results {
		if results[i] == nil {
			continue
		}
		for _, elem := range elements {
			if strings.HasPrefix(target[i:], elem) {
				combinations := make([][]string, 0, len(target)+1)
				if i == 0 {
					combinations = append(combinations, []string{elem})
				} else {
					for _, res := range results[i] {
						c := append(res, elem)
						combinations = append(combinations, c)
					}
				}
				results[len(elem)+i] = append(results[len(elem)+i], combinations...)
			}
		}
	}

	return results[len(target)]
}
