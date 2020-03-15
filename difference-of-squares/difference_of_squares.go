//Package diffsquares implements methods to calculate difference between square of the sum and sum of the suqares for first N natural numbers
package diffsquares

// SquareOfSum calculates square of sum for first n natural numbers
func SquareOfSum(n int) int {
	return (n * (n + 1) / 2) * (n * (n + 1) / 2)
}

// SumOfSquares calculates sum of squares for first n natural numbers
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference calculates difference between square of the sums and sum of the suqares for first n natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
