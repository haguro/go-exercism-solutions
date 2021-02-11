// Package letter provides functions to calculate the Frequency of letters in strings.
package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency concurrently calls `Frequency` on multiple strings and
// combines the results.
func ConcurrentFrequency(xs []string) FreqMap {
	fm := FreqMap{}
	c := make(chan FreqMap, len(xs))
	for _, s := range xs {
		go func(s string, c chan FreqMap) {
			c <- Frequency(s)
		}(s, c)

	}
	for range xs {
		for r, f := range <-c {
			fm[r] += f
		}
	}
	return fm
}
