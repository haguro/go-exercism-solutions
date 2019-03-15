//Package collatzconjecture provides bility to test the Collatz Conjecture
package collatzconjecture

import "errors"

//CollatzConjecture returns the number of steps required to reach one following the Collatz Conjecture
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n needs to be a positive integer")
	}
	steps := 0
	for n > 1 {
		switch n % 2 {
		case 0:
			n /= 2
		case 1:
			n = 3*n + 1
		}
		steps++
	}
	return steps, nil
}
