package main

func containsDuplicate(nums []int) bool {
	found := make(map[int]struct{}, len(nums))

	for _, n := range nums {
		if _, ok := found[n]; ok {
			return true
		}

		found[n] = struct{}{}
	}

	return false
}
