// Package isogram will check that given word is isogram
package isogram

import (
	"unicode"
)

// IsIsogram function will validate that letter does not repeat and resturn bool
func IsIsogram(s string) bool {
	a := map[rune]bool{}
	for _, r := range s {
		p := unicode.ToLower(r)
		if unicode.IsLetter(p) && a[p] {
			return false
		}
		a[p] = true
	}
	return true
}
