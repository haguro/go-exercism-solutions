//Package dna provides the ability to count nucleotide frequencies in a DNA strand.
package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

const bases = "ACGT"

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, n := range d {
		if !strings.ContainsRune(bases, n) {
			return nil, errors.New("DNA contains invalid nucleotide(s)")
		}
		h[n]++
	}
	return h, nil
}
