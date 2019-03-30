// Package hamming calculates Hamming distance
package hamming

import "errors"

// Distance calculates Hamming distance between two sequences
func Distance(a, b string) (int, error) {

	var res int

	if len(a) != len(b) {
		return res, errors.New("Sequences are not of equal length")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			res++
		}
	}

	return res, nil
}
