package main

import "fmt"

func main() {
	fmt.Println(fib(6))
	fmt.Println(fib(7))
	fmt.Println(fib(9))

	fb := make(fastFib)
	fmt.Println(fb.fib(50))

	fmt.Println(gridTraveler(1, 1))
	fmt.Println(gridTraveler(2, 3))
	fmt.Println(gridTraveler(3, 2))

	ft := make(memoGridTraveler)
	fmt.Println(ft.travel(18, 18))
}

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

type fastFib map[int]int

func (f fastFib) fib(n int) int {
	res, ok := f[n]
	if ok {
		return res
	}

	if n <= 2 {
		return 1
	}

	f[n] = f.fib(n-1) + f.fib(n-2)
	return f[n]
}

func gridTraveler(m, n int) int {
	if m == 1 && n == 1 {
		return 1
	}

	if m == 0 || n == 0 {
		return 0
	}

	return gridTraveler(m-1, n) + gridTraveler(m, n-1)
}

type memoGridTraveler map[string]int

func (t memoGridTraveler) travel(m, n int) int {
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

	t[key] = t.travel(m-1, n) + t.travel(m, n-1)
	return t[key]
}
