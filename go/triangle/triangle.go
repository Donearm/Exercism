package triangle

type Kind string

const (
	Equ = "equilateral"
	Iso = "isosceles"
	Sca = "scalene"
	NaT = "not a triangle"
)

func KindFromSides(a, b, c float64) Kind {
	switch {
	// Negative sides, NaN and Inf
	case !(a > 0 && b > 0 && c > 0):
		return NaT
	// Check for triangle inequality
	case (a+b) <= c || (a+c) <= b || (b+c) <= a:
		return NaT
	// All sides equal, equilateral
	case a == b && b == c:
		return Equ
	// 2 sides equal, isosceles
	case a == b || b == c || a == c:
		return Iso
	// Otherwise it must be a scalene
	default:
		return Sca
	}
}
