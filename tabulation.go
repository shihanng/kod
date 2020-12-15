package main

func fibTabulation(n int) int {
	fibNumbers := make([]int, n+1)
	fibNumbers[1] = 1

	for i := range fibNumbers {
		if i < n {
			fibNumbers[i+1] += fibNumbers[i]
		}

		if i < n-1 {
			fibNumbers[i+2] += fibNumbers[i]
		}
	}

	return fibNumbers[n]
}

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
