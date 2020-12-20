package gridtraveler

import "fmt"

func recursive(m, n int) int {
	if m == 1 && n == 1 {
		return 1
	}

	if m == 0 || n == 0 {
		return 0
	}

	return recursive(m-1, n) + recursive(m, n-1)
}

type memo map[string]int

func (t memo) recursive(m, n int) int {
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

	t[key] = t.recursive(m-1, n) + t.recursive(m, n-1)
	return t[key]
}

func tabulation(m, n int) int {
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
