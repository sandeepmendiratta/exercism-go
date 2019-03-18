// Package raindrops for converting a number to a string, depending upon on the number's factors
package raindrops

import (
	"strconv"
)

// Convert function is used for checking the factor and output string
func Convert(number int) string {
	var speech = []struct {
		talk string
		val  int
	}{
		{val: 3, talk: "Pling"},
		{val: 5, talk: "Plang"},
		{val: 7, talk: "Plong"},
	}
	var s string
	for _, k := range speech {
		if number%k.val == 0 {
			s += k.talk
		}
	}
	if s == "" {
		return strconv.Itoa(number)
	}
	return s
}
