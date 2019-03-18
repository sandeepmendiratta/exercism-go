// Package hamming for calculating hamming between two strings as interger
package hamming

import "fmt"

// Distance function will take two strings and will output hamming distance int
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("input strings %v and %v are not equal", a, b)
	}
	dist := 0
	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist, nil
}
