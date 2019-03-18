package luhn

import (
	"math"
	"strconv"
	"strings"
)

// Valid function is to calculate luhn algorithm
func Valid(num string) bool {
	var isValid bool
	var total float64
	var secondDigitTotal float64
	var eventDigit int
	num = strings.Replace(num, " ", "", -1)
	if len(num) > 1 {
		for i := len(num) - 1; i >= 0; i-- { //start from last num and go backwards
			number, error := strconv.Atoi(string(num[i]))
			// Check to make sure it is int otherwise exit loop
			if error != nil {
				return isValid
			}
			if eventDigit%2 != 0 { //check if it 2nd digit otherwise add num to total
				if number*2 > 9 {
					secondDigitTotal += float64(number)*2 - 9
				} else {
					secondDigitTotal += float64(number) * 2
				}
			} else {
				total += float64(number)
			}
			eventDigit++
		}
		total += secondDigitTotal
		if math.Mod(total, 10) == 0 {
			isValid = true
		}
	}
	return isValid
}
