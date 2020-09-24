package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("hamming: sequences have no equal length")
	}

	var hammingDistance int

	for i, ch := range a {
		if byte(ch) != b[i] {
			hammingDistance += 1
		}
	}

	return hammingDistance, nil
}
