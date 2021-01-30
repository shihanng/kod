package work_break

import "strings"

func wordBreak(s string, wordDict []string) bool {
	sList := make([]string, 0, len(s))
	sList = append(sList, s)
	memo := make(map[string][]string, len(sList))

	for len(sList) > 0 {
		var target string
		target, sList = sList[0], sList[1:]

		if target == "" {
			return true
		}

		if _, ok := memo[target]; ok {
			continue
		}

		var trimmeds []string
		for _, w := range wordDict {
			trimmed := strings.TrimPrefix(target, w)

			if len(trimmed) < len(target) {
				trimmeds = append(trimmeds, trimmed)
			}
		}

		memo[target] = trimmeds
		sList = append(sList, trimmeds...)
	}

	return false
}
