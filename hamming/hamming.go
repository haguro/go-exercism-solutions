// Package hamming provides the ability to calculate the Hamming difference between two DNA strands.
package hamming

import "errors"

// Distance returns the Hamming distance between two strand strings.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("squences have different lengths")
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
