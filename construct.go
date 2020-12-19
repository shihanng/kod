package main

import (
	"strings"
)

type canConstruct map[string]bool

func (c canConstruct) can(target string, elements []string) bool {
	if res, ok := c[target]; ok {
		return res
	}

	if target == "" {
		return true
	}

	for _, elem := range elements {
		if !strings.HasPrefix(target, elem) {
			continue
		}

		if c.can(strings.TrimPrefix(target, elem), elements) {
			c[target] = true
			return true
		}
	}

	c[target] = false
	return false
}

type countConstruct map[string]int

func (c countConstruct) count(target string, elements []string) int {
	if res, ok := c[target]; ok {
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

		if n := c.count(strings.TrimPrefix(target, elem), elements); n > 0 {
			count += n
		}
	}

	c[target] = count
	return count
}

type allConstruct map[string][][]string

func (c allConstruct) all(target string, elements []string) [][]string {
	if res, ok := c[target]; ok {
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

		res := c.all(strings.TrimPrefix(target, elem), elements)
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

	c[target] = all

	return all
}

func canConstructTabulation(target string, elements []string) bool {
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

func countConstructTabulation(target string, elements []string) int {
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

func allConstructTabulation(target string, elements []string) [][]string {
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
