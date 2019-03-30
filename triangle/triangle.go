// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
// Package triangle determines whether a triangle is Equilateral, isosceles, scalene type or is not a triangle.
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = "Not a triangle"
	Equ Kind = "Equilateral"
	Iso Kind = "Isosceles"
	Sca Kind = "Scalene"
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {

	var k Kind

	if (a > 0 && b > 0 && c > 0) && (a+b >= c && a+c >= b && b+c >= a) && !(math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1)) {
		// Check for Equ
		if a == b && b == c && a != 0 {
			k = Equ
		}
		// Check for Iso
		if (a == b && a != c) || (a == c && a != b) || (b == c && b != a) {
			k = Iso
		}
		// Check for Sca
		if a != b && a != c && b != c {
			k = Sca
		}
	} else {
		k = NaT
	}
	return k
}
