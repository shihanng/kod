package main

import "fmt"

func main() {
	fmt.Println(fib(6))
	fmt.Println(fib(7))
	fmt.Println(fib(9))
	fmt.Println(fib(50))
}

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
