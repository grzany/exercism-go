// Package grains implements functions to calculate how many grains on the chess board for a given square and total number
package grains

import "fmt"

const min uint64 = 1
const max uint64 = 64

//Square calculates how many grains for a given square
func Square(i int) (uint64, error) {
	// dismiss invalid values,i.e. square numbers outside of the chess board
	if i < int(min) || i > int(max) {
		return 0, fmt.Errorf("Invalid value %d", i)
	}
	//calculate 2^(i-1) which is a number of grains for square "i"
	return min << (i - 1), nil
}

//Total calculates total number of grains on the chess board
func Total() uint64 {
	// total number of grains is equal to sum of powers of 2 for n = 0 .. 63.  2^0 + 2^1 + ... 2^63
	// This is a geometric series which can be represented as 2^(n+1) - 1, for n = 63 it gives 2^64 -1
	return 1<<max - 1
}
