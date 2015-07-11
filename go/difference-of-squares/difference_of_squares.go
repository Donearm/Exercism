package diffsquares

func SquareOfSums(n int) int {
	var t int
	for x := 1; x <= n; x++ {
		t += x
	}
	return t*t
}

func SumOfSquares(n int) int {
	// make a slice to contain the squares of n
	a := make([]int, n)
	for x := 1; x <= n; x++ {
		a = append(a, x*x)
	}
	// add up all the squares
	var t int
	for _, v := range a {
		t += v
	}
	return t
}

func Difference(a int) int {
	return SquareOfSums(a) - SumOfSquares(a)
}
