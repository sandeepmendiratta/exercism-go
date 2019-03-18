// Package twofer provides string output
package twofer

// ShareWith function takes the sting input and provides string output
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
