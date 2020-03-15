//Package pythagorean implements functions to check certain properties of triangles
package pythagorean

//Triplet defines type to store a triangle
type Triplet [3]int

//Range returns a list of all Pythagorean triplets with sides in the range min to max
func Range(min, max int) []Triplet {
	var t []Triplet
	var a, b, c int
	for a = min; a <= max; a++ {
		for b = a + 1; b <= max; b++ {
			for c = b + 1; c <= max; c++ {
				if isPythagoreanTriangle(Triplet{a, b, c}) {
					t = append(t, Triplet{a, b, c})
				}
			}
		}
	}
	return t
}

//Sum returns a list of all Pythagorean triplets where sum is equal to p
func Sum(p int) []Triplet {
	var t []Triplet
	var a, b, c int
	sum := a + b + c

	for c = 3; sum <= p; c++ {

		for b = c - 1; b > 0; b-- {
			for a = b - 1; a > 0; a-- {
				if isPythagoreanTriangle(Triplet{a, b, c}) && a+b+c == p {
					var d []Triplet
					d = append(d, Triplet{a, b, c})
					// prepend triplet to t
					t = append(d, t...)
				}
			}
		}
		sum = a + b + c
	}
	return t
}

func isPythagoreanTriangle(t Triplet) bool {
	return t[0]*t[0]+t[1]*t[1] == t[2]*t[2]
}
