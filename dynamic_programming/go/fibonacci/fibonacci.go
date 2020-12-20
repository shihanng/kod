package fibonacci

func recursive(n int) int {
	if n <= 2 {
		return 1
	}
	return recursive(n-1) + recursive(n-2)
}

type memo map[int]int

func (f memo) recursive(n int) int {
	res, ok := f[n]
	if ok {
		return res
	}

	if n <= 2 {
		return 1
	}

	f[n] = f.recursive(n-1) + f.recursive(n-2)
	return f[n]
}

func tabulation(n int) int {
	sequence := make([]int, n+1)
	sequence[1] = 1

	for i := range sequence {
		if i < n {
			sequence[i+1] += sequence[i]
		}

		if i < n-1 {
			sequence[i+2] += sequence[i]
		}
	}

	return sequence[n]
}
