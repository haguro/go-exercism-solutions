// Package grains provides the ability to calculate the number of grains on a
// chess board where the number doubles on each square.
package grains

import (
	"errors"
)

// Square returns the number of grains for a given chessboard square.
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("square needs to be a positive integer between 1 and 64")
	}
	return 1 << (n - 1), nil
}

// Total returns the total number of grains on the chessboard.
func Total() uint64 {
	return (1 << 64) - 1 //The sum of powers of 2 up to 2^n is (2^n+1)-1
}
