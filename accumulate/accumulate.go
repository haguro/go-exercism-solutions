// Package accumulate provide functionality to apply a function
// on a given set strings.
package accumulate

// Accumulate returns a given slice of strings 'c' after applying
// a given function 'f' on them.
func Accumulate(c []string, f func(string) string) []string {
	var r = make([]string, len(c))
	for i, s := range c {
		r[i] = f(s)
	}
	return r
}
