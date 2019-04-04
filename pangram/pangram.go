// Package pangram provides the functionality to detect whether a
// sentence is a pangram.
package pangram

import (
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// IsPangram asserts whether a given string is a Pangram or not.
func IsPangram(s string) bool {
	s = strings.ToLower(s)
	for _, letter := range alphabet {
		if !strings.ContainsRune(s, letter) {
			return false
		}
	}
	return true
}
