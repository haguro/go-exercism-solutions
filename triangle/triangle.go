// Package triangle provides functionality to determine a triangle type
package triangle

import "math"

//Kind is a numerical representation of a triangle kind
type Kind int

//nolint
const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
	Deg        // degenerate
)

// KindFromSides returns triangle type based on given side lengths
func KindFromSides(a, b, c float64) Kind {
	switch {
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c):
		return NaT
	case math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1):
		return NaT
	case math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1):
		return NaT
	case a <= 0 || b <= 0 || c <= 0 || a+b < c || a+c < b || b+c < a:
		return NaT
	case a+b == c || a+c == b || b+c == a:
		return Deg
	case a == b && b == c:
		return Equ
	case (a == b && b != c) || (a == c && c != b) || (b == c && c != a):
		return Iso
	default:
		return Sca
	}
}
