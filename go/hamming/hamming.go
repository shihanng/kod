// Package hamming provides function to compute Hamming distance between two DNA strands.
package hamming

import "errors"

// Distance computes the Hamming distance between two DNA strands.
// The string inputs must have same length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("hamming: sequences have no equal length")
	}

	var hammingDistance int

	for i, ch := range a {
		if byte(ch) != b[i] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
