package prime

import (
	"math"
)

type primeAndMutiple struct {
	p, m int
}

// Nth returns the nth prime number
// Adapted from Mark Hollomon's implementation of the incremental sieve of
// Eratosthenes algorithm https://www.codevamping.com/2019/01/incremental-sieve-of-eratosthenes/
func Nth(n int) (int, bool) {
	switch {
	case n <= 0:
		return 0, false
	case n == 1:
		return 2, true
	}
	var pml []primeAndMutiple
	c, x := 1, 1
	for c < n {
		x += 2 //All primes that are larger than 2 are odd
		limit := int(math.Sqrt(float64(x)))
		isPrime := true
		for _, pm := range pml {
			if pm.p > limit {
				break
			}
			for pm.m < x {
				pm.m += pm.p
			}
			if pm.m == x {
				isPrime = false
				break
			}
		}
		if isPrime {
			pml = append(pml, primeAndMutiple{x, x})
			c++
		}
	}
	return x, true
}
