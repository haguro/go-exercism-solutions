// Package acronym provides the ability to generate acronyms from multi-word phrases.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate returns a string of capitalised first letters from each word.
func Abbreviate(s string) string {
	var abbr []byte
	words := regexp.MustCompile(`[\w']+`).FindAll([]byte(s), -1)
	for _, word := range words {
		abbr = append(abbr, word[0])
	}
	return strings.ToUpper(string(abbr))
}
