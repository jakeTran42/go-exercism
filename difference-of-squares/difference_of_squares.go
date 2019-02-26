package diffsquares

// SquareOfSum : Inputs an integer firstN and returns the square of
// the sum of the first N natural numbers
func SquareOfSum(firstN int) int {
	sum := firstN * (firstN + 1) / 2
	square := sum * sum
	return square
}

// SumOfSquares : Inputs an integer firstN and returns the sum of the
// squares of the first N natural numbers
func SumOfSquares(firstN int) int {
	sum := firstN * (firstN + 1) * (2*firstN + 1) / 6
	return sum
}

// Difference : Inputs an integer firstN and returns a difference
// between the square of sums and the sum of squares
func Difference(firstN int) int {
	return SquareOfSum(firstN) - SumOfSquares(firstN)
}
