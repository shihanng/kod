package main

func gridTravellerTabulation(m, n int) int {
	grid := make([][]int, n+1)
	for j := 0; j <= n; j++ {
		grid[j] = make([]int, m+1)
	}

	grid[1][1] = 1

	for j := 0; j <= n; j++ {
		for i := 0; i <= m; i++ {
			if i != m {
				grid[j][i+1] += grid[j][i]
			}
			if j != n {
				grid[j+1][i] += grid[j][i]
			}
		}
	}

	return grid[n][m]
}

func canSumTabulation(targetSum int, numbers []int) bool {
	targets := make([]bool, targetSum+1)
	targets[0] = true

	for i := range targets {
		if targets[i] {
			for _, n := range numbers {
				if i+n <= targetSum {
					targets[i+n] = true
				}
			}
		}
	}

	return targets[targetSum]
}

func howSumTabulation(targetSum int, numbers []int) []int {
	targets := make([][]int, targetSum+1)

	targets[0] = []int{}

	for i := range targets {
		if targets[i] != nil {
			for _, n := range numbers {
				if i+n <= targetSum && targets[i+n] == nil {
					targets[i+n] = append(targets[i], n)
				}
			}
		}
	}

	return targets[targetSum]
}

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
