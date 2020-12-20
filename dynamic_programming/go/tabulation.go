package main

func bestSumTabulation(targetSum int, numbers []int) []int {
	targets := make([][]int, targetSum+1)

	targets[0] = []int{}

	for i := range targets {
		if targets[i] != nil {
			for _, n := range numbers {
				newSum := append(targets[i], n)
				if i+n <= targetSum && (targets[i+n] == nil || len(newSum) < len(targets[i+n])) {
					targets[i+n] = newSum
				}
			}
		}
	}

	return targets[targetSum]
}
