// Package diffsquares provides functions to calculate sum of squares and
// squared sum for a given series of natural numbers.
package diffsquares

// SquareOfSum calcuates the sum of a series of natural numbers (1 to n)
// using Faulhaber's formula (where p = 1) then squares it.
func SquareOfSum(n int) int {
	return n * n * (n + 1) * (n + 1) / 4
}

// SumOfSquares calcuates the sum of squares of a series of natural numbers
// (1 to n) using Faulhaber's formula (where p=2).
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns the difference between the square of the sum of a series
// of natural numbers (1 to n) and the sum of their squares.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
