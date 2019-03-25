//Package strain provides ability to keep or discard items in collections
//based a specific predicate.
package strain

//Ints represents a collection of integers
type Ints []int

//Lists represents a collection of `Ints`
type Lists [][]int

//Strings represents a collection of strings
type Strings []string

//Keep returns an Ints collection of items (int) kept by function f.
func (ints Ints) Keep(f func(int) bool) Ints {
	var kept Ints
	for _, i := range ints {
		if f(i) {
			kept = append(kept, i)
		}
	}
	return kept
}

//Discard returns an Ints collection of items (int) discarded by function f.
func (ints Ints) Discard(f func(int) bool) Ints {
	var discarded Ints
	for _, i := range ints {
		if !f(i) {
			discarded = append(discarded, i)
		}
	}
	return discarded
}

//Keep returns a Lists collection of items (Ints) kept by function f.
func (lists Lists) Keep(f func([]int) bool) Lists {
	var kept Lists
	for _, l := range lists {
		if f(l) {
			kept = append(kept, l)
		}
	}
	return kept
}

//Keep returns a Strings collection of items (string) kept by function f.
func (strings Strings) Keep(f func(string) bool) Strings {
	var kept Strings
	for _, s := range strings {
		if f(s) {
			kept = append(kept, s)
		}
	}
	return kept
}
