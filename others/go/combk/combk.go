package combk

func Combk(n, k int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	all := [][]int{{n}}

	for i := 1; i < n; i++ {
		combinations := Combk(i, k)
		for _, c := range combinations {
			if len(c) < k {
				all = append(all, append(c, n))
			} else {
				all = append(all, c)
			}
		}
	}
	all = append(all, []int{n})

	return all
}
