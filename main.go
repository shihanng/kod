package main

import "fmt"

func main() {
	fmt.Println(fib(6))
	fmt.Println(fib(7))
	fmt.Println(fib(9))

	fb := make(fastFib)
	fmt.Println(fb.fib(50))
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
