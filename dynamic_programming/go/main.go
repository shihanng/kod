package main

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
