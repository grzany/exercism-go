// Package dna provides way to count number of nucleotides in a DNA
package dna

import "errors"

// Histogram keeps counts of nucleotides in a DNA
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for i, _ := range d {
		_, exists := h[rune(d[i])]
		if !exists {
			return h, errors.New("Invalid nucleotide")
		}
		h[rune(d[i])]++
	}
	return h, nil
}
