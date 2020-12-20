package main

import (
	"fmt"
)

func main() {
	fmt.Println(gridTraveler(1, 1))
	fmt.Println(gridTraveler(2, 3))
	fmt.Println(gridTraveler(3, 2))

	ft := make(memoGridTraveler)
	fmt.Println(ft.travel(18, 18))
}

func gridTraveler(m, n int) int {
	if m == 1 && n == 1 {
		return 1
	}

	if m == 0 || n == 0 {
		return 0
	}

	return gridTraveler(m-1, n) + gridTraveler(m, n-1)
}

type memoGridTraveler map[string]int

func (t memoGridTraveler) travel(m, n int) int {
	key := fmt.Sprintf("%d,%d", m, n)

	val, ok := t[key]
	if ok {
		return val
	}

	if m == 1 && n == 1 {
		return 1
	}

	if m == 0 || n == 0 {
		return 0
	}

	t[key] = t.travel(m-1, n) + t.travel(m, n-1)
	return t[key]
}

func canSum(targetSum int, numbers []int) bool {
	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	for _, n := range numbers {
		remainder := targetSum - n
		if canSum(remainder, numbers) {
			return true
		}
	}

	return false
}

type fastCanSum map[int]bool

func (c fastCanSum) canSum(targetSum int, numbers []int) bool {
	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	if res, ok := c[targetSum]; ok {
		return res
	}

	for _, n := range numbers {
		remainder := targetSum - n
		c[targetSum] = c.canSum(remainder, numbers)
		if c[targetSum] {
			return true
		}
	}

	c[targetSum] = false

	return false
}

func howSum(targetSum int, numbers []int) []int {
	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	for _, n := range numbers {
		remainder := targetSum - n

		res := howSum(remainder, numbers)
		if res != nil {
			return append(res, n)
		}
	}

	return nil
}

type fastHowSum map[int][]int

func (s fastHowSum) howSum(targetSum int, numbers []int) []int {
	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	if res, ok := s[targetSum]; ok {
		return res
	}

	for _, n := range numbers {
		remainder := targetSum - n

		s[targetSum] = s.howSum(remainder, numbers)
		if s[targetSum] != nil {
			s[targetSum] = append(s[targetSum], n)
			return s[targetSum]
		}
	}

	s[targetSum] = nil
	return s[targetSum]
}

type bestSum map[int][]int

func (s bestSum) bestSum(targetSum int, numbers []int) []int {
	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	if res, ok := s[targetSum]; ok {
		return res
	}

	var shortest []int

	for _, n := range numbers {
		remainder := targetSum - n

		res := s.bestSum(remainder, numbers)
		if res == nil {
			continue
		}

		res = append(res, n)

		if len(shortest) == 0 || len(res) < len(shortest) {
			shortest = res
		}
	}

	s[targetSum] = shortest
	return shortest
}
