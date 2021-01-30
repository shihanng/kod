package dup

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums)-1; i++ {
		var subNums []int
		j := i + k + 1

		if j <= len(nums) {
			subNums = nums[i+1 : j]
		} else {
			subNums = nums[i+1:]
		}

		for _, n := range subNums {
			if abs(nums[i]-n) <= t {
				return true
			}

		}

	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
