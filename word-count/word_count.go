// Package wordcount provides functionality to count word
// occurances in a given string.
package wordcount

import (
	"regexp"
	"strings"
)

// Frequency represent a map of words and their number of occurances.
type Frequency map[string]int

var rx = regexp.MustCompile(`[[:alnum:]]+('\w)?`)

// WordCount returns the frequency of words in a given sentence.
func WordCount(s string) Frequency {
	var f = make(Frequency)
	words := rx.FindAllString(strings.ToLower(s), -1)
	for _, w := range words {
		f[w]++
	}
	return f
}
