package leap

import "fmt"

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}
func main() {
	fmt.Println(IsLeapYear(2015))
}
