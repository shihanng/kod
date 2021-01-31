package rank

import "sort"

func arrayRankTransform(arr []int) []int {
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)

	sort.Ints(arrCopy)

	ranks := make(map[int]int)

	currentRank := 1
	for _, n := range arrCopy {
		if _, ok := ranks[n]; ok {
			continue
		}

		ranks[n] = currentRank
		currentRank++
	}

	for i, n := range arr {
		arrCopy[i] = ranks[n]
	}

	return arrCopy
}
