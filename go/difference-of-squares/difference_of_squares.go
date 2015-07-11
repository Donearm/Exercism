package diffsquares

func SquareOfSums(n int) int {
	var t int
	for x := 1; x <= n; x++ {
		t += x
	}
	return t*t
}

func SumOfSquares(n int) int {
	var t int
	for x := 1; x <= n; x++ {
		t += x*x
	}
	return t
}

func Difference(a int) int {
	return SquareOfSums(a) - SumOfSquares(a)
}
