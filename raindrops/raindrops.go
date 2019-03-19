//Package raindrops provides the functionalty to generate raindrop-speak strings.
package raindrops

import "strconv"

//Convert returns the "raindrop-speak" version of a given number.
func Convert(n int) string {
	s := ""
	if n%3 == 0 {
		s += "Pling"
	}
	if n%5 == 0 {
		s += "Plang"
	}
	if n%7 == 0 {
		s += "Plong"
	}
	if s == "" {
		return strconv.Itoa(n)
	}
	return s
}
