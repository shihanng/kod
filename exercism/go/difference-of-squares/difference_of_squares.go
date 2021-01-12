// Package diffsquares provides function to compute "Difference Of Squares"
package diffsquares

// SquareOfSum computes the square of the sum of the first n natural numbers.
func SquareOfSum(n int) int {
	sum := n * (1 + n) / 2

	return sum * sum
}

// SumOfSquares computes the sum of the squares of the first n natural numbers.
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference is SquareOfSum - SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
