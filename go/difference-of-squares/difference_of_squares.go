package diffsquares

// SquareOfSum function will to add sum of number and square
func SquareOfSum(num int) int {
	sum := (num * (num + 1) / 2)
	return sum * sum
}

// SumOfSquares function will square and sum
func SumOfSquares(num int) int {
	sum := (num * (num + 1) * (2*num + 1)) / 6
	return sum

}

// Difference function will give Difference between SquareOfSum and SumOfSquares
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)

}
