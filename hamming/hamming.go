package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var error error
	if len(a) != len(b) {
		error = errors.New("Length of String a and b not equal")
		return 0, error
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
