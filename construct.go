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
