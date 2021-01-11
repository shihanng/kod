// Package diffsquares provides function to compute "Difference Of Squares"
package diffsquares

// SquareOfSum computes the square of the sum of the first ten natural numbers.
func SquareOfSum(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum * sum
}

// SumOfSquares computes the sum of the squares of the first ten natural numbers.
func SumOfSquares(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i * i
	}

	return sum
}

// Difference is SquareOfSum - SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
