package combk

import "fmt"

func Combk(n, k int) [][]int {
	combinations := make(n, k)

	ans := [][]int{}
	found := map[string]bool{}

	for _, c := range combinations {
		if len(c) != k {
			continue
		}

		h := fmt.Sprintf("%v", c)
		if found[h] {
			continue
		}

		ans = append(ans, c)
		found[h] = true
	}

	return ans
}

func make(n, k int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	all := [][]int{{n}}

	for i := 1; i < n; i++ {
		combinations := make(i, k)
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
