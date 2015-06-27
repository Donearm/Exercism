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
		case (a + b) <= c || (a + c) <= b || (b + c) <= a: return NaT
		case (a == b) && (b == c): return Equ
		case (a == b) || (b == c) || (a == c): return Iso
		case (a + b) > c || (a + c) > b || (b + c) > a: return Sca
		default: return NaT
	}
}
