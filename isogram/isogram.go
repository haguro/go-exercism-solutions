// Package isogram provides the functionality to test whether a word is an isogram or not.
package isogram

import "unicode"

// IsIsogram returns `true` if the given word is an isogram, `false` if it is not.
func IsIsogram(word string) bool {
	seen := make(map[rune]bool)
	for _, r := range word {
		if r == ' ' || r == '-' {
			continue
		}
		r = unicode.ToLower(r)
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}
