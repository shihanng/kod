package main

import "strings"

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
