// Package listops povides serveral useful list operations.
package listops

// IntList represents a list of integers.
type IntList []int

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Foldr recurively combines (folds) the values in a given list
// (from right to left) to a single value based ona given function.
func (list IntList) Foldr(f binFunc, init int) int {
	r := init
	for i := len(list); i > 0; i-- {
		r = f(list[i-1], r)
	}
	return r
}

// Foldl recurively combines (folds) the values in a given list
// (from left to right) to a single value based ona given function.
func (list IntList) Foldl(f binFunc, init int) int {
	r := init
	for i := 0; i < len(list); i++ {
		r = f(r, list[i])
	}
	return r
}

// Filter returns the values from a given list based on output of
// a given function.
func (list IntList) Filter(f predFunc) IntList {
	var r = IntList{}
	for i := 0; i < len(list); i++ {
		if f(list[i]) {
			r = append(r, list[i])
		}
	}
	return r
}

// Length return the value count of a given list
func (list IntList) Length() int {
	return len(list)
}

// Map returns list of value resulting from applying a given function
// to the elements of a given list.
func (list IntList) Map(f unaryFunc) IntList {
	var r = IntList{}
	for i := 0; i < len(list); i++ {
		r = append(r, f(list[i]))
	}
	return r
}

// Reverse returnes a version of a given list with element order reversed.
func (list IntList) Reverse() IntList {
	var r = IntList{}
	for i := len(list) - 1; i >= 0; i-- {
		r = append(r, list[i])
	}
	return r
}

// Append combines two list into one.
func (list IntList) Append(appended IntList) IntList {
	r := list
	for i := 0; i < len(appended); i++ {
		r = append(r, appended[i])
	}
	return r
}

// Concat combines multiple list into one.
func (list IntList) Concat(args []IntList) IntList {
	r := list
	for i := 0; i < len(args); i++ {
		r = r.Append(args[i])
	}
	return r
}
