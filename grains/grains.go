package grains

import (
	"errors"
	"math"
)

// Square : Square inputs a day, represented as an integer
// between 1 and 64 inclusive, and returns how much grain
// should be recieved on that day, as an integer.
func Square(squareNum int) (uint64, error) {
	if squareNum < 1 || squareNum > 64 {
		return 0, errors.New("Input day invalid")
	}

	grains := math.Pow(2, float64(squareNum-1))
	return uint64(grains), nil
}

// Total : Total returns the total number of grains recieved
// after 64 days
func Total() uint64 {
	return 18446744073709551615
}
