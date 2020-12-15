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
