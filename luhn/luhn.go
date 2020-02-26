// Package luhn provides the functionality to determine whether a given number
// is valid per the Luhn formula.
package luhn

import (
	"strconv"
	"strings"
)

// Valid returns `true` if the given number is valid.
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	l := len(s)
	if l <= 1 {
		return false
	}
	var sum int
	for i := l - 1; i >= 0; i-- {
		n, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return false
		}
		if (l-i)%2 == 0 {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
	}
	return sum%10 == 0
}
